from flask import Flask, request, escape
import re
import urllib.parse

app = Flask(__name__)


@app.route("/")
def root():
    args = request.args

    address = "https://startpage.com"

    if "q" in args:
        query = args["q"]

        # Exclamation point, any amount of none whitespace characters, either a whitespace character or the end of the line.
        regex = re.search("![^\s]*(\s|$)", query)

        address = (
            f"https://startpage.com/do/metasearch.pl?query={urllib.parse.quote(query)}"
        )

        if regex:
            address = f"https://duckduckgo.com/?q={urllib.parse.quote(query)}"

    return f"<html><head><meta http-equiv='Content-Type' content='text/html; charset=utf-8'><meta name='referrer' content='never'><meta name='robots' content='noindex, nofollow'><meta http-equiv='refresh' content='0; url={escape(address)}'></head><body><script language='JavaScript'>function ffredirect(){{window.location.replace('{escape(address)}');}}setTimeout('ffredirect()',100);</script></body></html>"


if __name__ == "__main__":
    # Listen on all interfaces, this can be insecure however I'm a pro-gamer!
    app.run(host="0.0.0.0", port=3000)
