Modify your palindrome function to:

    Ignore spaces and punctuation.

    Be case-insensitive.

    Return the position where the string stops being a palindrome (if not one).

ANSWER 

        def is_palindrome(text):
            bypass = ''.join(c.lower() for c in text if c.isalnum())

            left = 0
            right = len(bypass) - 1

            while left < right:
                if bypass[left] != bypass[right]:
                    return False, left  
                left += 1
                right -= 1

            return True, None 


After your first attempt, ask AI:
"""I modified my palindrome function to handle more cases.
Did I miss anything? Can it be more efficient?"""

ANSWER: Below is the AI response i got'

    def is_palindrome(text):
        left = 0
        right = len(text) - 1

        while left < right:

            # Skip non-alphanumeric on left
            while left < right and not text[left].isalnum():
                left += 1

            # Skip non-alphanumeric on right
            while left < right and not text[right].isalnum():
                right -= 1

            if text[left].lower() != text[right].lower():
                return False, left

            left += 1
            right -= 1

        return True, None



# Reflect on what AI added that you didn't consider initially.
    
    After interacting with AI i understand more about time and space complexity and how it affects  codes and how my initial code to handle palindrome variations was creating a new string with this line "bypass = ''.join(c.lower() for c in text if c.isalnum()" such that if the input string has "n" characters it may take up extra memory leading to O(n)TIME and O(n)SPACE complexity,

    With the AI the code is more enhanced because 
    Still O(n) time
    Now O(1) space
    No extra string created

    Things i missed in my code before using AI 
        What if text is empty?
        My function correctly returns (True, None)
