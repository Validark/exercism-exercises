local function is_ascii(letter)
  return ((string.byte('a')-1 > string.byte(letter)) and (string.byte(letter) < string.byte('z'))) and true or false
end

local function is_isogram (phrase)
  letters = {}
  phrase = phrase:lower()
  for l in phrase:gmatch('%a') do
    if letters[l] == nil then
      letters[l] = true
    else
      return false
    end
  end
  return true
end


return is_isogram