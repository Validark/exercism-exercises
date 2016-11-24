local function count (phrase)
  local split = {}
  for word in string.gmatch(phrase:lower(), '%w+') do
     split[word] = (split[word] or 0) + 1
  end
  return split
end

return {
  word_count = count
}