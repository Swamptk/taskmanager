package db

func FilterUndone(tasks []Task) []Task {
	ret := []Task{}
	for _, t := range tasks {
		if !t.Value.Done {
			ret = append(ret, t)
		}
	}
	return ret
}

func FilterDone(tasks []Task) []Task {
	ret := []Task{}
	for _, t := range tasks {
		if t.Value.Done {
			ret = append(ret, t)
		}
	}
	return ret
}
