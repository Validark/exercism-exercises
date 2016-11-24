local bracket = {}

function bracket.valid(str)
  local stack = {}
  for i=1, #str do
    rune = str:sub(i,i)
    if rune == '{' or rune == '(' or rune == '[' then
       table.insert(stack, rune)
    elseif rune == '}' or rune == ')' or rune == ']'  then
       local x = table.remove(stack)
       if rune == '}' and x ~= '{' then
          return false
      elseif rune == ')' and x ~= '(' then
         return false
      elseif rune == ']' and x ~= '[' then
          return false
      end
    end
  end
  if #stack ~=0 or #stack % 2 ~= 0 then
     return false
  end
  return true
end

return bracket