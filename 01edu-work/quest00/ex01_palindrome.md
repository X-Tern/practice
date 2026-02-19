
**Pseudocode for a function that checks if a string is a palindrome.**
	
	psuedocode for the palindrome:

                START
            .Algorithm is_palindrome
            .input string(text)
            .output is boolean ("True" if palindrome and "False" if otherwise)
            .initialize left set to zero
            .initialize right TO length of text minus one

            .WHILE left is lessthan right
            .If characters of text on [left] is NOT equal to text on [right]
            .RETURN False
            .END IF

            .move left pointer by 1
            .move right pointer backward by 1
            .END WHILE

            .RETURN True
            END


**Python inplementation of the solution.**

    def is_palindrome(text):

         # Start one pointer at the beginning of the string
        left = 0

        # Start another pointer at the end of the string
        right = len(text) - 1


         # Keep checking characters while the two pointers
        # have not crossed each other
        while left < right:


        # Compare the character at the left pointer
        # with the character at the right pointer
            if text[left] != text[right]:

            # If they are not equal, the string is NOT a palindrome
                return False

            # Move the left pointer one step to the right and 
            left += 1 

            # Move the right pointer one step to the left
            right -= 1

    # If the loop finishes without finding a mismatch,
    # it means every compared pair matched
    # Therefore, the string IS a palindrome

    return True




# STEP 3.

    (1).QUESTION: What did you learn from solving it before asking AI?
    ANSWER: While solving it on my own i learnt that writing a code for palindrome i needed to add two pointers one from left and one from the right of the string which will compare the characters involed to be sure they are same if not terminate immediatly.

    (2) QUESTION: How is your understanding different now?
        ANSWER: After using AI i got to understand that it isn't just about checking the characters in the string if they are thesame but also i need to clean the string if it behaves differently or contains other characters like spaces between words, also found out a shorter way to write the code as shown below.

        def is_palindrome(text):
            return text == text[::-1]

        print(is_palindrome("radar"))
        print(is_palindrome("Level"))

    
    (3) QUESTION: Could you now write similar functions (e.g., reverse a string) without help?
        ANSWER: Yes i can write similar functions.

        def is_palindrome(text):
            return text[::-1]

        print(is_palindrome("radar"))
        print(is_palindrome("Level"))
