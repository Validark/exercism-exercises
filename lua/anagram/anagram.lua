local Anagram = {}
Anagram.__index = Anagram


local function sortword(word)
  local w = {}
  word = word:lower()
  for i=1, #word do
    w[i] = word:sub(i,i)
  end
  table.sort(w)
  return table.concat(w, '')
end

function Anagram:new(word)
  local o = {}
  setmetatable(o, self)
  o.word = word:lower()
  return o
end

function Anagram:match(list)
  local match = {}
  for k, w in ipairs(list) do
    if sortword(w) == sortword(self.word) then
      table.insert(match, w)
    end
  end
  return match
end

return Anagram