package practice

// Function to find the longest common prefix among a list of strings
func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0]

    for i := 1; i < len(strs); i++ {
        // Compare the prefix with the current string
        prefix = commonPrefix(prefix, strs[i])
        // If the common prefix becomes an empty string, no need to check further
        if prefix == "" {
            break
        }
    }

    return prefix
}

// Helper function to calculate the common prefix of two strings
func commonPrefix(str1, str2 string) string {
    minLen := len(str1)
    if len(str2) < minLen {
        minLen = len(str2)
    }

    for i := 0; i < minLen; i++ {
        if str1[i] != str2[i] {
            return str1[:i]
        }
    }

    return str1[:minLen]
}
