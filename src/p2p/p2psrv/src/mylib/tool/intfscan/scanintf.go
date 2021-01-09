package intfscan

type ScanIntf interface {
	SetAggregate(AggregateIntf)
	Run()
	RunOnce()
}
