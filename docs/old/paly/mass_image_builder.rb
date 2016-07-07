
File.open("html.txt","w") do |f|
	51.times do |i|
		f << "<p><img  width=\"320\" src=\"320/#{i}_320.jpg\" ></p> \n"
	end
end