import flask
import requests

app = flask.Flask(__name__)

@app.route("/route_param/<route_param>")
def route_param(route_param):
    print("blah")
    # ruleid: ssrf-requests
    return requests.get(route_param)

@app.route("/route_param_ok/<route_param>")
def route_param_ok(route_param):
    print("blah")
    # ok: ssrf-requests
    return requests.get("this is safe")

@app.get("/route_param/<route_param>")
def route_param_without_decorator(route_param):
    print("blah")
    # ruleid: ssrf-requests
    return requests.get(route_param)

@app.route("/get_param", methods=["GET"])
def get_param():
    param = flask.request.args.get("param")
    # ruleid: ssrf-requests
    requests.post(param, timeout=10)
