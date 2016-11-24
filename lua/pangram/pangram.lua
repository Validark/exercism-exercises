local function is_pangram (phrase)
  letters = {}
  for i in string.gmatch(phrase, '%w') do
    i = i:lower()
    letters[i] = (letters[i] or 0) + 1
  end
  if #letters == 26 then
    return true
  else
    return false
  end
end

return