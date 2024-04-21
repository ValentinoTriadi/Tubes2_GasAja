package main

// Utils function to check if a web entity is already in the slice
func containsWebEntity (webEntity web, Res []web) bool {
	for _, w := range Res {
		if w.Url == webEntity.Url {
			return true
		}
	}
	return false
}

func isStorageContainsWebEntity(webEntity web, storage *[][]ResultEntity) bool {
	for _, res := range *storage {
		for _, r := range res {
			if r.webEntity.Url == webEntity.Url {
				return true
			}
		}
	}
	return false
}

func appendToResult(storage *[][]ResultEntity, level int, index int, webEntity web, saveRes *[][]web) {
	// Get the result
	var result = []web{}

	// Append the web entity to the result
	result = append(result, webEntity)

	// Get the parent index
	var parentIndex = index

	// Loop through the storage to get the result
	for i := level; i >= 0; i-- {
		// Append the web entity to the result
		result = append(result, (*storage)[i][parentIndex].webEntity)
		// Get the parent index
		parentIndex = (*storage)[i][parentIndex].index
	}

	// Reverse the result
	reverse(&result)

	// Append the result to the saveRes
	if isResultNotInSaveRes(result, saveRes) {
		*saveRes = append(*saveRes, result)
	}
}

func isResultNotInSaveRes(result []web, saveRes *[][]web) bool {
	for _, res := range *saveRes {
		if isSameResult(res, result) {
			return false
		}
	}
	return true
}

func isSameResult(res1 []web, res2 []web) bool {
	if len(res1) != len(res2) {
		return false
	}
	for i := 0; i < len(res1); i++ {
		if res1[i].Url != res2[i].Url {
			return false
		}
	}
	return true
}


func reverse(res *[]web) {
	for i, j := 0, len(*res)-1; i < j; i, j = i+1, j-1 {
		(*res)[i], (*res)[j] = (*res)[j], (*res)[i]
	}
}