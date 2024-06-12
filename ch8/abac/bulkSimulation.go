//given Main function
func (ss *Sim) main(){
    ss.New()
	ss.Config()

	ss.Init()
	win := ss.ConfigGui()
	win.StartEventLoop()

    ss.newTrain()
}

//train for 50 runs 
func (ss *Sim) newTrain() {

	ss.Init()
	i := 1

    type set func(ss *Sim, x float64)
    set := setInhib
    parameter int := Inhibition
    start float64 := 0;
    end float64 := 0;
    incr float64 := 0;
    numSeeds int := 1;
    numTraining int := 5


    switch parameter; {
    case Inhibition:
        set := setInhib
    case InitialWeight:
        set := setInitWt
    case Context:
        set := setContext
    case XCal:
        set := setXCAL
    case LearningRate:
        set := Lrate
    }

    for k := 0; k < numSeeds; k++ {
        ss.newRndSeed()
        for i := start; i <= end; i += incr {
            set(ss, i)
            ss.init()
            for j := 0; j < numTraining; j++ {
                ss.Train()
            }
        }
    }
}

//parameter types
const (

    Inhibition int = Itoa
    InitialWeight
    Context
    XCal
    LearningRate
)


//set functions
func (ss *Sim, x float64) setInhib{
	ss.HiddenInhibGi = x
}

func (ss *Sim, x float64) setInitWt{
	ss.WtInitVar = x
}

func (ss *Sim, x float64) setContext{
	ss.FmContext = x
}

func (ss *Sim, x float64) setXCAL{
	ss.XCalLLrn = x
}

func (ss *Sim, x float64) setLrate{
	ss.Lrate = x
}
