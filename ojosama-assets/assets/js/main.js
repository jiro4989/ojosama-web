async function convert() {
  const text = document.getElementById("input").value;
  if (text === "") {
    return;
  }

  document.getElementById("progress").innerHTML = "処理中";
  const data = {"Text": text}
  const url = "https://api.ojosama.jiro4989.com/"
  const resp = await fetch(url, {
    method: "POST",
    mode: 'cors',
    headers: {
      "Content-type": "application/json"
    },
    body: JSON.stringify(data)
  });
  const result = await resp.json();
  document.getElementById("output").value = result.Result;
  progress.innerHTML = "変換完了！";

  activeButton("tweet-button", result.Result !== "");
}

function tweet() {
  const url = location.protocol + "//" + location.host;
  const output = document.getElementById("output").value;
  const message = `${output}\n\nOjosama web converter\n`;
  const encodedMessage = encodeURIComponent(message);
  const encodedURL = encodeURIComponent(url);
  const tweetURL = "https://twitter.com/intent/tweet" + `?url=${encodedURL}&text=${encodedMessage}`;
  window.open(tweetURL, "");
}

function activeButton(elemId, active) {
  const button = document.getElementById(elemId);
  if (active === true) {
    button.removeAttribute("disabled");
  } else {
    button.setAttribute("disabled", true);
  }
}

function changeConvertButtonState() {
  const text = document.getElementById("input").value;
  activeButton("convert-button", text !== "");
}
