function raindrops (num)
    local str = {}
    for k,v in pairs(factor(num)) do
        if v == 3 then
            table.insert(str, 'Pling')
        elseif v == 5 then
            table.insert(str, 'Plang')
        elseif v == 7 then
            table.insert(str, 'Plong')
        end
    end
    return #str ~= 0 and table.concat(str, '') or tostring(num)
end

function factor(num)
    f = {}
    for i = 0, num do
        fact = num % i
        if fact == 0 then
            table.insert(f, i)
        end
    end
    return f
end

return raindrops