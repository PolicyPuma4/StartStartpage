import express from "express";

const app = express();

app.get("/", (req, res) => {
  const query = req.query.q;

  let address = "https://startpage.com";
  if (query) {
    const regex = new RegExp("!([^\\s]+)", "i");

    address = `https://startpage.com/do/metasearch.pl?query=${encodeURIComponent(
      query
    )}`;
    if (regex.test(query)) {
      address = `https://duckduckgo.com/?q=${encodeURIComponent(query)}`;
    }
  }

  res.send(
    `<html><head><meta http-equiv='Content-Type' content='text/html; charset=utf-8'><meta name='referrer' content='never'><meta name='robots' content='noindex, nofollow'><meta http-equiv='refresh' content='0; url=${address}'></head><body><script language='JavaScript'>function ffredirect(){window.location.replace('${address}');}setTimeout('ffredirect()',100);</script></body></html>`
  );
});

app.listen(3000, () => {
  console.log("Listening on port", 3000);
});

