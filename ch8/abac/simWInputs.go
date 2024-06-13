
package main

import "fmt"
//given Main function
func (ss *Sim) main(){
    ss.New()
	ss.Config()

	ss.Init()
	win := ss.ConfigGui()
	win.StartEventLoop()

    mode := 0
    fmt.Println("Enter Mode:")
    fmt.Println("(1)Use default parameters for all but the parameter being trained")
    fmt.Println("(2)Accumulate changes to each parameter in training")
    fmt.Scanf("%d", &mode)
    for mode != 1 && mode != 2 {
        fmt.Println("Enter Mode 1 or 2:")
        fmt.Println("(1)Use default parameters for all but the parameter being trained")
        fmt.Println("(2)Accumulate changes to each parameter in training")
        fmt.Scanf("%d", &mode)
    }


    amountParams := 0
    fmt.Println("How many parameters do you want to train?")
    fmt.Scanf("%d", &amountParams)
    for !(amountParams >= 1 && amountParams <= 5) {
        fmt.Println("Enter integer between 1 and 5:")
        fmt.Scanf("%d", &amountParams)
    }

    paramOrder := []Parameter{}
    var param Parameter
    startValues := []float32{}
    endValues := []float32{}
    increments := []float32{}

    for i := 0; i < amountParams; i++ {

        

        fmt.Println("Choose a parameter to alter:")
        fmt.Println("(1)Inhibition  (2)Initial Weights  (3)Context  (4)XCAL Learning  (5)Learning Rate")
        fmt.Scanf("%d", &param)
        for !(param >= 1 && param <= 5)  {
            fmt.Println("Enter valid parameter between 1 and 5:")
            fmt.Println("(1)Inhibition  (2)Initial Weights  (3)Context  (4)XCAL Learning  (5)Learning Rate")
            fmt.Scanf("%d", &param)
        }

        paramOrder = append(paramOrder, param)

        startValues = append(startValues, startInput())
        endValues = append(endValues, endInput(startValues[i]))
        increments = append(increments, incrInput())


    }

    numTraining := 0
    fmt.Printf("How many times do you want to train for each value?\n")
    fmt.Scanf("%d", &numTraining)
    for numTraining <= 0 {
        fmt.Printf("Choose training amount value greater than 0\n")
        fmt.Scanf("%d", &numTraining)
    }

    fmt.Printf("Ready to Train\n")

    

    var setParam func(ss *Sim, x float32)
    setParam = setInhib

    for i := 0; i < amountParams; i++ {

        if mode == 1 {
            ss.Defaults()
        }

        param = paramOrder[i]
        switch; {
        case param == Inhibition:
            setParam = setInhib
        case param == InitialWeight:
            setParam = setInitWt
        case param == Context:
            setParam = setContext
        case param == XCal:
            setParam = setXCAL
        case param == LearningRate:
            setParam = setLrate
        }

        fmt.Printf("Training parameter %s on interval %f - %f %d times at each %f increment\n", paramOrder[i].String(), startValues[i], endValues[i], numTraining, increments[i])

        ss.newTrain(startValues[i], endValues[i], increments[i], numTraining, setParam)


    }

    


    




}

//input version
func (ss *Sim) newTrain(startVal float32, endVal float32, incr float32, numTraining int, set func(ss *Sim, x float32)) {


    for i := startVal; i <= endVal; i += incr {
        set(ss, i)
        ss.NewRndSeed()
        ss.Init()
        for j := 0; j < numTraining; j++ {
            ss.Train()
        }
    }
    
}

type Parameter int
//parameter types
const (

    Inhibition Parameter = iota + 1
    InitialWeight
    Context
    XCal
    LearningRate
)

func (p Parameter) String() string{
    return []string{"Inhibition", "Initial Weight", "Context", "XCal Learning", "Learning Rate"}[p - 1]
}

func (p Parameter) enumInd() int {
    return int(p)
}


//set functions
func setInhib(ss *Sim, x float32){
	ss.HiddenInhibGi = x
}

func setInitWt(ss *Sim, x float32){
	ss.WtInitVar = x
}

func setContext(ss *Sim, x float32){
	ss.FmContext = x
}

func setXCAL(ss *Sim, x float32){
	ss.XCalLLrn = x
}

func setLrate(ss *Sim, x float32){
	ss.Lrate = x
}


func startInput() float32 {

    var startVal float32 = 0
    fmt.Printf("Choose a start value:\n")
    fmt.Scanf("%f", &startVal)
    for startVal < 0 {
        fmt.Printf("Choose start value greater than or equal to 0\n")
        fmt.Scanf("%f", &startVal)
    }

    return startVal
}


func endInput(start float32) float32 {

    var endVal float32 = 0
    fmt.Printf("Choose an end value:\n")
    fmt.Scanf("%f", &endVal)
    for endVal < start {
        fmt.Printf("Choose end value greater than start\n")
        fmt.Scanf("%f", &endVal)
    }

    return endVal
}


func incrInput() float32 {

    var incr float32 = 0
    fmt.Printf("Choose an increment value:\n")
    fmt.Scanf("%f", &incr)
    for incr <= 0 {
        fmt.Printf("Choose increment value greater than 0\n")
        fmt.Scanf("%f", &incr)
    }

    return incr
}
