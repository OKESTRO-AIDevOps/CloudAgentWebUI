FROM ubuntu:20.04

RUN apt update

RUN apt install -y openjdk-11-jre openjdk-11-jdk wget gnupg2 curl

RUN wget -q -O - https://pkg.jenkins.io/debian-stable/jenkins.io.key | apt-key add -

RUN sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'

RUN apt update

RUN apt install -y jenkins

RUN apt install -y apt-transport-https

RUN apt install -y software-properties-common wget

RUN wget -q -O /usr/share/keyrings/grafana.key https://packages.grafana.com/gpg.key

RUN echo "deb [signed-by=/usr/share/keyrings/grafana.key] https://packages.grafana.com/enterprise/deb stable main" | tee -a /etc/apt/sources.list.d/grafana.list

RUN apt update

RUN apt install -y grafana-enterprise

RUN wget https://github.com/vmware-tanzu/octant/releases/download/v0.24.0/octant_0.24.0_Linux-64bit.deb

RUN dpkg -i octant_0.24.0_Linux-64bit.deb

RUN curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

RUN chmod +x /usr/local/bin/docker-compose

RUN apt update

RUN apt install -y apt-transport-https ca-certificates curl

RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list

RUN apt update

RUN apt install -y ca-certificates curl gnupg lsb-release

RUN mkdir -p /etc/apt/keyrings

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt update

RUN apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

RUN apt update

RUN apt install -y kubectl

RUN curl https://baltocdn.com/helm/signing.asc | apt-key add - 

RUN apt install apt-transport-https --yes 

RUN echo "deb https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list 

RUN apt update 

RUN apt install -y helm

RUN apt install -y vim

RUN apt install -y nginx

RUN apt update

RUN apt -y install python3-pip

RUN pip3 install flask

RUN pip3 install pandas

RUN pip3 install python-dotenv

RUN pip3 install requests

RUN pip3 install gunicorn


EXPOSE 80
EXPOSE 8080
EXPOSE 3000
EXPOSE 7777

VOLUME /var/run/docker.sock
VOLUME /root/.jenkins
VOLUME /var/lib/grafana

CMD ["tail","-f","/dev/null"]
