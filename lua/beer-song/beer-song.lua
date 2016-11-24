local beer = {}

local function combine(list)
  local l = {}
  for i=1, #list do
    if i ~= #list then table.insert(l, list[i] .. '\n') else table.insert(l, list[i]) end
  end
  return table.concat(l)
end

function beer.verse(x)
    if x > 2 then
      return string.format("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n",x, x, x-1)
    elseif x > 1 then
        return string.format("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n",x, x, x-1)
    elseif x == 1 then
      return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
    elseif x == 0 then
      return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
    else return ''
    end
end

function beer.sing(from, to)
  local song = {}
  to = to or 0
  for i=from, to, -1 do
    table.insert(song, beer.verse(i))
  end
  return combine(song)
end

return beer