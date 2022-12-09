package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func octant_run() {

	docker := "docker"
	execute := "exec"
	d := "-d"
	cake8s := "cake8s"
	binsh := "/bin/sh"
	c := "-c"
	octant := "OCTANT_DISABLE_OPEN_BROWSER=true OCTANT_LISTENER_ADDR=0.0.0.0:7777 octant"

	fmt.Println("Initiating Octant to proceed...")

	cmd := exec.Command(docker, execute, d, cake8s, binsh, c, octant)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	fmt.Println("Octant initiated")

	fmt.Println("*** Octant ------> Running on localhost:7777")

	time.Sleep(3 * time.Second)
}

func jenkins_run() {

	docker := "docker"
	execute := "exec"
	copy := "cp"
	d := "-d"
	cake8s := "cake8s"
	binsh := "/bin/sh"
	c := "-c"
	jenkins := "jenkins"
	t := "-t"
	cat := "cat"
	initAdmPass := "/root/.jenkins/secrets/initialAdminPassword"
	jenkinsinitpass := "./ADM/jenkins_init_pass.txt"
	initiatedSign := "/root/.jenkins/jenkins_init_pass.txt"
	cake8sinitiatedSign := "cake8s:/root/.jenkins/jenkins_init_pass.txt"

	cmd := exec.Command(docker, execute, d, cake8s, binsh, c, jenkins)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	JINIT := 0

	for JINIT == 0 {

		var err1 error
		var err2 error

		cmd := exec.Command(docker, execute, t, cake8s, cat, initAdmPass)

		stdout, err1 := cmd.Output()

		cmd = exec.Command(docker, execute, t, cake8s, cat, initiatedSign)

		_, err2 = cmd.Output()

		if err1 != nil && err2 != nil {

			fmt.Println("Initiating Jenkins in process, please wait...")
			time.Sleep(5 * time.Second)

		} else if err1 == nil && err2 != nil {

			fmt.Println("Jenkins initiated")

			f, _ := os.Create("./ADM/jenkins_init_pass.txt")
			l, _ := f.WriteString(string(stdout))
			fmt.Println("*** Jenkins ------> Running on localhost:8080")
			fmt.Println(l, " bytes: INIT PASSWORD > ./ADM/jenkins_init_pass.txt")
			_ = f.Close()

			cmd := exec.Command(docker, copy, jenkinsinitpass, cake8sinitiatedSign)

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			time.Sleep(3 * time.Second)

			JINIT = 1

		} else {

			fmt.Println("Jenkins starting")
			fmt.Println("*** Jenkins ------> Running on localhost:8080")
			JINIT = 1
		}

	}

}

func grafana_run() {
	docker := "docker"
	execute := "exec"
	d := "-d"
	cake8s := "cake8s"
	binsh := "/bin/sh"
	c := "-c"
	grafana_server := "service grafana-server start"
	t := "-t"
	helm := "helm"
	repo := "repo"
	add := "add"
	prom_comm := "prometheus-community"
	prom_addr := "https://prometheus-community.github.io/helm-charts"
	//stable := "stable"
	//stable_addr := "https://charts.helm.sh/stable"
	update := "update"
	install := "install"
	prometheus := "prometheus"
	prom_name := "prometheus-community/prometheus"
	kubegetpods := "kubectl get pods --field-selector status.phase!=Running"
	portforwardrunner := "/home/runner.sh"

	cmd := exec.Command(docker, execute, d, cake8s, binsh, c, grafana_server)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, helm, repo, add, prom_comm, prom_addr)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	//cmd = exec.Command(docker, execute, t, cake8s, helm, repo, add, stable, stable_addr)

	//cmd.Stdout = os.Stdout

	//cmd.Stderr = os.Stderr

	//cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, helm, repo, update)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, helm, install, prometheus, prom_name)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	GINIT := 0

	for GINIT == 0 {

		cmd = exec.Command(docker, execute, t, cake8s, binsh, c, kubegetpods)

		stdout, _ := cmd.Output()

		stdout_str := string(stdout)

		if strings.Contains(stdout_str, "No resources") != true {

			fmt.Println("Initiating Grafana in process, please wait...")
			time.Sleep(5 * time.Second)

		} else if strings.Contains(stdout_str, "No resources") == true {

			cmd = exec.Command(docker, execute, d, cake8s, binsh, c, portforwardrunner)

			cmd.Stdout = os.Stdout

			cmd.Stderr = os.Stderr

			cmd.Run()

			fmt.Println("Grafana initiated")

			fmt.Println("*** Grafana ------> Running on localhost:3000")
			time.Sleep(3 * time.Second)

			GINIT = 1

		}

	}

}

func up(onWhat string) {

	if onWhat == "cake" {
		up_cake()
	} else if onWhat == "cake-reporter" {
		up_cake_reporter()
	}

}

func down() {

	down_cake()

}

func up_cake() {

	docker_compose := "docker-compose"
	up := "up"
	d := "-d"
	docker := "docker"
	copy := "cp"
	var ADMconfig string
	rootkubeconfig := "cake8s:/root/.kube/config"
	execute := "exec"
	cake8s := "cake8s"
	t := "-t"

	rm := "rm"
	mkdir := "mkdir"

	grafanaini := "./ADM/plugins/grafana.ini"
	pathgrafanaini := "cake8s:/etc/grafana/grafana.ini"
	etcgrafanaini := "/etc/grafana/grafana.ini"
	rootkube := "/root/.kube"
	hometarget := "/home/target"
	runnersh := "./ADM/plugins/runner.sh"
	pathrunnersh := "cake8s:/home/runner.sh"

	chmod := "chmod"
	x := "+x"
	homerunnersh := "/home/runner.sh"

	os.Chdir("./cake8s")

	fmt.Println("Initiating Octant, Jenkins, Grafana in that order...")

	fmt.Println("If you don't have the image already installed, it may take a while to set up...")

	time.Sleep(3 * time.Second)

	cmd := exec.Command(docker_compose, up, d)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	fmt.Println("KUBE CONFIG LOCATION : ")

	fmt.Scanln(&ADMconfig)

	cmd = exec.Command(docker, execute, t, cake8s, rm, etcgrafanaini)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, grafanaini, pathgrafanaini)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, mkdir, rootkube)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, mkdir, hometarget)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, runnersh, pathrunnersh)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, chmod, x, homerunnersh)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, ADMconfig, rootkubeconfig)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	octant_run()

	jenkins_run()

	grafana_run()

	fmt.Println("          ")
	fmt.Println("OPENING BROWSER")

	landing_url_octant := "http://localhost:7777"
	landing_url_jenkins := "http://localhost:8080"
	landing_url_grafana := "http://localhost:3000"

	if runtime.GOOS == "windows" {

		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", landing_url_octant).Start()

		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", landing_url_jenkins).Start()

		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", landing_url_grafana).Start()

	} else if runtime.GOOS == "linux" {

		_ = exec.Command("xdg-open", landing_url_octant).Start()

		_ = exec.Command("xdg-open", landing_url_jenkins).Start()

		_ = exec.Command("xdg-open", landing_url_grafana).Start()

	} else {

		fmt.Println("It looks like the browser is not a supported version or does not exist ...")

	}

	fmt.Println("         ")
	time.Sleep(3 * time.Second)

	fmt.Println("***")
	fmt.Println("*** Octant ------> Running on localhost:7777")
	fmt.Println("*** Jenkins ------> Running on localhost:8080")
	fmt.Println("***     Jenkins Init Password > ./ADM/jenkins_init_pass.txt")
	fmt.Println("*** Grafana ------> Running on localhost:3000")
	fmt.Println("***")
	fmt.Println("CAKE8S initiated successfully")
	fmt.Println("Use [ lwcli down ] to turn off")

}

func up_cake_reporter() {

	docker_compose := "docker-compose"
	up := "up"
	d := "-d"
	docker := "docker"
	copy := "cp"
	var ADMconfig string
	rootkubeconfig := "cake8s:/root/.kube/config"
	nginx := "service nginx start"
	execute := "exec"
	cake8s := "cake8s"
	binsh := "/bin/sh"
	c := "-c"
	t := "-t"

	gunicorn := "gunicorn -w 2 -b 0.0.0.0:7331 --chdir /home/reporter wsgi:app"
	rm := "rm"
	mkdir := "mkdir"
	pathavailable := "/etc/nginx/sites-available/default"
	pathenabled := "/etc/nginx/sites-enabled/default"
	defaultconf := "./ADM/plugins/reporter/default.conf"
	pathdefaultconf := "cake8s:/etc/nginx/conf.d/default.conf"
	grafanaini := "./ADM/plugins/grafana.ini"
	pathgrafanaini := "cake8s:/etc/grafana/grafana.ini"
	etcgrafanaini := "/etc/grafana/grafana.ini"
	rootkube := "/root/.kube"
	hometarget := "/home/target"
	runnersh := "./ADM/plugins/runner.sh"
	pathrunnersh := "cake8s:/home/runner.sh"
	reporter := "./ADM/plugins/reporter/."
	pathreporter := "cake8s:/home/reporter/"
	homereporter := "/home/reporter"
	chmod := "chmod"
	x := "+x"
	homerunnersh := "/home/runner.sh"

	os.Chdir("./cake8s")

	fmt.Println("Initiating Octant, Jenkins, Grafana in that order...")

	fmt.Println("If you don't have the image already installed, it may take a while to set up...")

	time.Sleep(3 * time.Second)

	cmd := exec.Command(docker_compose, up, d)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	fmt.Println("KUBE CONFIG LOCATION : ")

	fmt.Scanln(&ADMconfig)

	cmd = exec.Command(docker, execute, t, cake8s, rm, pathavailable)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, rm, pathenabled)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, defaultconf, pathdefaultconf)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, rm, etcgrafanaini)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, grafanaini, pathgrafanaini)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, mkdir, rootkube)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, mkdir, hometarget)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, runnersh, pathrunnersh)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, mkdir, homereporter)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, reporter, pathreporter)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, t, cake8s, chmod, x, homerunnersh)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, copy, ADMconfig, rootkubeconfig)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, d, cake8s, binsh, c, gunicorn)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	cmd = exec.Command(docker, execute, d, cake8s, binsh, c, nginx)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	octant_run()

	jenkins_run()

	grafana_run()

	fmt.Println("          ")
	fmt.Println("OPENING BROWSER")

	landing_url := "http://localhost"

	if runtime.GOOS == "windows" {

		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", landing_url).Start()

	} else if runtime.GOOS == "linux" {

		_ = exec.Command("xdg-open", landing_url).Start()

	} else {

		fmt.Println("It looks like the browser is not a supported version or does not exist ...")

	}

	fmt.Println("         ")
	time.Sleep(3 * time.Second)

	fmt.Println("***")
	fmt.Println("*** Jenkins Init Password > ./ADM/jenkins_init_pass.txt")
	fmt.Println("***")
	fmt.Println("CAKE8S initiated successfully")
	fmt.Println("Use [ lwcli down ] to turn off")

}

func down_cake() {

	docker_compose := "docker-compose"

	down := "down"

	os.Chdir("./cake8s")

	fmt.Println("Turning off Octant, Jenkins, Grafana...")

	cmd := exec.Command(docker_compose, down)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

}

func main() {

	action := os.Args[1]
	onWhat := os.Args[2]

	if action == "up" {

		up(onWhat)

	} else if action == "down" {

		down()

	}

}
