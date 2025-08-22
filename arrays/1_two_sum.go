package arrays

func TwoSum(nums []int, target int) []int {
    //HASHMAP
    m := make(map[int]int)
    for i, num := range nums {
        diff := target - num
        if idx, ok := m[diff]; ok {
            return []int{idx, i}
        }
        m[num] = i
    }
    return nil
    
    // BASED ON BRUTE FORCE
    // result := []int{}
    // for i := 0; i < len(nums) - 1; i++ {
    //     for j := i + 1; j < len(nums); j++ {
    //         if nums[i] + nums[j] == target {
    //             result = append(result, i)
    //             result = append(result, j)
    //             return result
    //         }
    //     }
    // }
    // return nil
}
