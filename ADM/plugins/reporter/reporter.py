from flask import Flask, request as request_flask, render_template
import subprocess
import json
from pathlib import Path
import os
import pandas as pd
import secrets
import requests
from dotenv import load_dotenv,find_dotenv

app = Flask(__name__)
port = 7331

def getApiAddr():

    out = subprocess.run(["kubectl","config","view","--minify","-o","jsonpath='{.clusters[].cluster.server}'"],capture_output=True)

    kubeapiaddr = out.stdout.decode('utf-8')

    return kubeapiaddr

def getAllResources():

    out = subprocess.run(["kubectl","get","all","--all-namespaces"],capture_output=True)

    kubeallresources = out.stdout.decode('utf-8')

    return kubeallresources

def getEvents():

    out = subprocess.run(["kubectl","get","events","--all-namespaces"],capture_output=True)

    kubeevents = out.stdout.decode('utf-8')

    return kubeevents


@app.route('/report', methods=['GET'])
def npiareport():

    post_dict = {"kubeapiaddr":'',"kubeallresources":'',"kubeevents":''}

    kubeapiaddr = getApiAddr()

    kubeallresources = getAllResources()

    kubeevents = getEvents()

    post_dict["kubeapiaddr"] = str(kubeapiaddr)

    post_dict["kubeallresources"] = str(kubeallresources)

    post_dict["kubeevents"] = str(kubeevents)

    return post_dict 



@app.route('/', methods=['GET'])
def npia():


    return render_template('index.html')


if __name__ == "__main__":
    app.run(host='0.0.0.0',port=port)