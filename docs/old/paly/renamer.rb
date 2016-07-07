#mass file renamer
dir = "__DIR__"
f = Dir["*.jpg"]
p f
i=1
f.each do |f|
	File.rename f , "#{i}.jpg"
	i +=1
end

