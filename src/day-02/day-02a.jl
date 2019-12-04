function atoi_array!(array::Array{SubString{String},1})
    array = map(array) do array
           parse(Int, array)
    end
    return array
end

s = open("input.txt") do file
    read(file, String)
end

# Parses the input into an array of numbers
memory = split(s, r",|\n|\r")
deleteat!(memory, length(memory)) # Last element after splitting the input is ""
memory = atoi_array!(memory)

# Adds the inputs to the Intcode program
memory[2] = 12 # Position 1, the noun
memory[3] = 2 # Position 2, the verb

# Runs the Intcode program
ip = 1
while (ip + 3) <= length(memory)
    if memory[ip] == 1
        memory[memory[ip+3] + 1] = memory[memory[ip+1] + 1] + memory[memory[ip+2] + 1]
    elseif memory[ip] == 2
        memory[memory[ip+3] + 1] = memory[memory[ip+1] + 1] * memory[memory[ip+2] + 1]
    elseif memory[ip] == 99
        println("Intcode program finished successfully.\nMemory after computation:\n")
        @show memory
        break
    else
        println("Intcode program encountered an error, halting execution...")
        break
    end
    global ip = ip + 4
end
