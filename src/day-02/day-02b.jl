function atoi_array!(array::Array{SubString{String},1})
    array = map(array) do array
           parse(Int, array)
    end
    return array
end

function findoutput(initialmemory::Array{Int64,1}, desiredoutput::Int64)
    for noun = 0:99
        for verb = 0:99
            memory = copy(initialmemory)

            # Adds the inputs to the Intcode program
            memory[2] = noun # Position 1, the noun
            memory[3] = verb # Position 2, the verb

            # Runs the Intcode program
            global ip = 1
            while (ip + 3) <= length(memory)
                if memory[ip] == 1
                    memory[memory[ip+3] + 1] = memory[memory[ip+1] + 1] + memory[memory[ip+2] + 1]
                elseif memory[ip] == 2
                    memory[memory[ip+3] + 1] = memory[memory[ip+1] + 1] * memory[memory[ip+2] + 1]
                elseif memory[ip] == 99
                    if memory[1] == desiredoutput # If ater the program ending, the output is the one desired
                        return noun, verb # Exits the function
                    else
                        global ip = length(memory) + 1 # If not, changes the IP to an invalid value so another iteration can start
                    end
                else
                    println("Intcode program encountered an error, halting execution...")
                    break
                end
                global ip = ip + 4
            end
        end
    end
end

s = open("input.txt") do file
    read(file, String)
end

# Parses the input into an array of numbers
initialmemory = split(s, r",|\n|\r")
deleteat!(initialmemory, length(initialmemory)) # Last element after splitting the input is ""
initialmemory = atoi_array!(initialmemory)

desiredoutput = 19690720
noun, verb = findoutput(initialmemory, desiredoutput)

println("Intcode program finished with the desired output.\nNoun and verb used:\n")
@show noun
@show verb
