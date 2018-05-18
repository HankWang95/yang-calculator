time = ['','21','22','23','24','01','02','03','04','05','06','07','08','09','10','11','12','13','14','15','16','17','18','19','20']

with open("p.txt", "rt",encoding="utf-8") as in_file_p:
	with open("down.txt", "wt",encoding="utf-8") as out_file:
		with open("wendu.txt", "rt",encoding="utf-8") as in_file_t:
			text = []
			text_list_two = []
			text_list_two_t = []

			for line in in_file_t:
				text.append(line)
				text_list_two_t.append(line.split(' '))



			for line in in_file_p:
				text.append(line)
				text_list_two.append(line.split(' '))

			            # text = text.replace("-", "* ")

			for i in range(len(text_list_two)):
				for x in range(len(text_list_two[i])-1):
					if text_list_two_t[i][x+1] == "////":
						continue
					tmp = int(text_list_two_t[i][x+1])
					tmp = tmp*0.1+273.16
					line = text_list_two[i][0]+" "+time[x+1]+" "+text_list_two[i][x+1]+" "+str(tmp)+"\r\n"
					out_file.write(line)