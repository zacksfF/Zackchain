package tracker

//CreateTracker a tracker instance by its well-known name
func createTracker(name string) ([]Tracker, error){
	switch name {
	case "timeout":
		{
			return []Tracker{&Time}
		}
		
	}
}