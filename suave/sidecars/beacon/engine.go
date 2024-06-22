package beacon_sidecar

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

type BeaconSidecar struct {
	latestBeaconBuildBlockArgs BeaconBuildBlockArgs
	latestTimestamp            uint64
	mu                         sync.Mutex
	cancel                     context.CancelFunc
	wg                         sync.WaitGroup
}

func NewBeaconSidecar(beaconRpc string, boostRelayUrl string) *BeaconSidecar {
	ctx, cancel := context.WithCancel(context.Background())
	sidecar := &BeaconSidecar{
		cancel: cancel,
	}
	go sidecar.startSyncing(ctx, beaconRpc, boostRelayUrl)
	return sidecar
}

func (bs *BeaconSidecar) GetLatestBeaconBuildBlockArgs() BeaconBuildBlockArgs {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.latestBeaconBuildBlockArgs.Copy()
}

func (bs *BeaconSidecar) GetLatestTimestamp() uint64 {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.latestTimestamp
}

func (bs *BeaconSidecar) Stop() {
	bs.cancel()
	bs.wg.Wait()
}

func (bs *BeaconSidecar) startSyncing(ctx context.Context, beaconRpc string, boostRelayUrl string) {
	defer bs.wg.Done()
	defer bs.cancel()

	bs.wg.Add(1)
	payloadAttrC := make(chan PayloadAttributesEvent)
	go SubscribeToPayloadAttributesEvents(ctx, beaconRpc, payloadAttrC)

	for paEvent := range payloadAttrC {

		log.Debug("New PA event", "data", paEvent.Data)

		validatorData, err := getValidatorForSlot(ctx, boostRelayUrl, paEvent.Data.ProposalSlot)
		if err != nil {
			log.Warn("could not get validator", "slot", paEvent.Data.ProposalSlot, "err", err)
			continue
		}

		res, err := json.Marshal(validatorData)
		if err != nil {
			log.Error("could not marshal validator data", "err", err)
			continue
		}
		log.Debug("New validator data", "data", string(res))

		var beaconBuildBlockArgs = BeaconBuildBlockArgs{
			Slot:            paEvent.Data.ProposalSlot,
			ProposerPubkey:  hexutil.MustDecode(validatorData.Pubkey),
			Parent:          paEvent.Data.ParentBlockHash,
			Timestamp:       paEvent.Data.PayloadAttributes.Timestamp,
			Random:          paEvent.Data.PayloadAttributes.PrevRandao,
			FeeRecipient:    validatorData.FeeRecipient,
			GasLimit:        validatorData.GasLimit,
			ParentBlockRoot: paEvent.Data.PayloadAttributes.ParentBeaconBlockRoot,
		}

		for _, w := range paEvent.Data.PayloadAttributes.Withdrawals {
			withdrawal := Withdrawal{
				Index:     uint64(w.Index),
				Validator: uint64(w.ValidatorIndex),
				Address:   common.Address(w.Address),
				Amount:    uint64(w.Amount),
			}
			beaconBuildBlockArgs.Withdrawals = append(beaconBuildBlockArgs.Withdrawals, withdrawal)
		}

		bs.update(beaconBuildBlockArgs)
	}
}

func (bs *BeaconSidecar) update(beaconBuildBlockArgs BeaconBuildBlockArgs) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.latestBeaconBuildBlockArgs = beaconBuildBlockArgs
	bs.latestTimestamp = uint64(time.Now().Unix())
}

type BeaconBuildBlockArgs struct {
	Slot            uint64
	ProposerPubkey  []byte
	Parent          common.Hash
	Timestamp       uint64
	FeeRecipient    common.Address
	GasLimit        uint64
	Random          common.Hash
	Withdrawals     []Withdrawal
	ParentBlockRoot common.Hash
}

func (b *BeaconBuildBlockArgs) Copy() BeaconBuildBlockArgs {
	deepCopy := BeaconBuildBlockArgs{
		Slot:            b.Slot,
		ProposerPubkey:  make([]byte, len(b.ProposerPubkey)),
		Parent:          b.Parent,
		Timestamp:       b.Timestamp,
		Random:          b.Random,
		FeeRecipient:    b.FeeRecipient,
		GasLimit:        b.GasLimit,
		ParentBlockRoot: b.ParentBlockRoot,
		Withdrawals:     make([]Withdrawal, len(b.Withdrawals)),
	}

	copy(b.ProposerPubkey, deepCopy.ProposerPubkey)
	for i, w := range b.Withdrawals {
		deepCopy.Withdrawals[i] = Withdrawal{
			Index:     w.Index,
			Validator: w.Validator,
			Address:   w.Address,
			Amount:    w.Amount,
		}
	}

	return deepCopy
}

func (b BeaconBuildBlockArgs) Bytes() ([]byte, error) {
	return json.Marshal(b)
}

type Withdrawal struct {
	Index     uint64
	Validator uint64
	Address   common.Address
	Amount    uint64
}
