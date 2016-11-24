diff = {}

function squared(x) return x * x end
function val(x) return x end

table.reduce = function (list, fn)
    if #list == 0 then
       return 0
    end
    local acc
    for k, v in ipairs(list) do
        if 1 == k then
            acc = v
        else
            acc = fn(acc, v)
        end
    end
    return acc
end

local function accumulate (num, f)
  ret = {}
  for i = 1, num  do
    table.insert(ret,f(i))
  end
  return ret
end

local function sum (arr)
  return table.reduce(
      arr,
      function (a, b)
          return a + b
      end
  )
end

local function diff.square_of_sums(num)
  return sum(accumulate(num, val))^2
end


local function diff.sum_of_squares(num)
    return sum(accumulate(num, squared))
end

function diff.difference_of_squares(num)
    print(diff.sum_of_squares(num))
    print(diff.square_of_sums(num))
    return diff.square_of_sums(num) - diff.sum_of_squares(num)
end

return diff