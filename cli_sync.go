package main

func (cli *CLI) sync() {
	currentIp := getIPV4()
	status := "syn" + string(currentIp)
	send_status("13.40.159.39", status)
	recv_file()
}
