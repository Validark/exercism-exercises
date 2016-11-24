hamming = {}

function hamming.compute(seq1, seq2)
    local count = 0
    for i=1,  #seq1 do
        local s1 = seq1:sub(i,i)
        local s2 = seq2:sub(i,i)
        if s1 ~= s2 then
            count = count + 1
        end
    end
    return count
end
return hamming