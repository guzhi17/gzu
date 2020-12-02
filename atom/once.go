package atom


//this once not used to wait for initialization, it is use to close something once
type Once struct {
	done Int32
}
func (o *Once) Do(f func()) {
	if o.done.Inc() == 1{
		f()
	}
}
func (o *Once) IsDone()bool{
	return 0!=o.done.Load()
}

func (o *Once) Reset(){
	o.done.Store(0)
}
