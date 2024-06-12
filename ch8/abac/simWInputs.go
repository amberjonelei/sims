//given Main function
func (ss *Sim) main(){
    ss.New()
	ss.Config()

	ss.Init()
	win := ss.ConfigGui()
	win.StartEventLoop()

    ss.newTrain()
}

//input version
func (ss *Sim) newTrain() {

	ss.Init()
	i := 1

    parameter int := Inhibition
    int := x
    fmt.printf("Choose a variable to alter: \n")
    fmt.printf("(1)Inhibition  (2)Initial Weights  (3)Context  (4)XCAL Learning  (5)Learning Rate\n")
    fmt.scanf("%d", &x)

    parameter = x - 1
    type set func(ss *Sim, x float64)
    set := setInhib

    start float64 := 0;
    end float64 := 0;
    incr float64 := 0;
    numSeeds int := 1;
    numTraining int := 5

    fmt.printf("Choose a start value:\n")
    fmt.scanf("%f", &start)
    for start < 0 {
        fmt.printf("Choose start value greater than or equal to 0\n")
        fmt.scanf("%f", &start)
    }

    fmt.printf("Choose an end value:\n")
    fmt.scanf("%f", &end)
    for end < start {
        fmt.printf("Choose end value greater than start\n")
        fmt.scanf("%f", &end)
    }

    fmt.printf("Choose an increment value:\n")
    fmt.scanf("%f", &incr)
    for incr <= 0 {
        fmt.printf("Choose increment value greater than 0\n")
        fmt.scanf("%f", &incr)
    }

    fmt.printf("How many times do you want to train for each value?\n")
    fmt.scanf("%d", &numTraining)
    for numTraining <= 0 {
        fmt.printf("Choose training amount value greater than 0\n")
        fmt.scanf("%d", &incr)
    }

    fmt.printf("How many seeds do you want to use?\n")
    fmt.scanf("%d", &seeds)
    for numTraining <= 0 {
        fmt.printf("Choose seed amount value greater than 0\n")
        fmt.scanf("%d", &incr)
    }

    fmt.printf("Ready to Train\n")

    //which parameter
    //range and increment

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

