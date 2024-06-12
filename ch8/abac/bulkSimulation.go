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

//chris psuedo
func run_simulation(model, data, parms) {
    model.New()
    // pass the parameters here or somewhere else to set them
    // maybe we can use a hashmap of parameters so that it is variable length
    // and the parameter values can be binded the model object via keyword?
    model.Config(parms)//doesnt take params use set
    // formerly CmdArgs
    model.setupAndRun()
    data.appendData(model)
}

func main(){

    // initialize model
    // initialize parameters
    // initialize data 
    // loop through parameters
    for parm in parms { 
        run_simulation(model, data, parms)    
    }
    return data 
}