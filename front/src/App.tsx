import React, { useEffect, useState, ChangeEvent } from "react";
import { QRCodeServiceClient } from "./proto/qrcode_pb_service";
import { Empty, QRCode, URL } from "./proto/qrcode_pb";
import { BrowserHeaders } from "browser-headers";
import "./App.css";

const App = () => {
  const [data, setData] = useState<{ items: QRCode.AsObject[] }>({ items: [] });
  const [title, setTitle] = useState<string>("");
  const [pdfURL, setPDFURL] = useState<string>("");

  const baseURL = location.protocol + "//" + location.host + "/grpc-web";
  const client = new QRCodeServiceClient(baseURL);

  const fetchData = async () => {
    client.getQRCodes(new Empty(), new BrowserHeaders(), (e, res) => {
      if (e || !res) {
        throw e;
      } else {
        const list = res.toObject().qrcodesList;
        setData({ items: list });
      }
    });
    client.getURL(new Empty(), new BrowserHeaders(), (e, res) => {
      if (e || !res) {
        throw e;
      } else {
        setPDFURL(res.toObject().url);
      }
    });
  };

  const handleChangeTitle = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setTitle(e.target.value);
  };
  const handleChangePDFURL = (e: ChangeEvent<HTMLInputElement>) => {
    setPDFURL(e.target.value);
  };

  const onClickAddQRCode = async () => {
    client.addQRCodes(new Empty(), new BrowserHeaders(), () => {});
  };
  const onClickSetPDFURL = async () => {
    const req = new URL();
    req.setUrl(pdfURL);
    client.updateURL(req, new BrowserHeaders(), () => {});
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="App">
      <div className="uk-margin no_print">
        <input
          className="uk-input"
          type="text"
          placeholder="PDFのURL"
          value={pdfURL}
          onChange={handleChangePDFURL}
        />
        <textarea
          className="uk-textarea"
          rows={3}
          placeholder="タイトル"
          value={title}
          onChange={handleChangeTitle}
        />
        <button className="uk-button uk-button-default" onClick={onClickSetPDFURL}>
          URL変更
        </button>
        <button className="uk-button uk-button-default" onClick={onClickAddQRCode}>
          QRコード追加
        </button>
      </div>
      <div className="panel">
        {data.items
          .reduce(
            (table: QRCode.AsObject[][], item) => {
              const last = table[table.length - 1];
              if (last.length === 2) {
                table.push([item]);
                return table;
              }
              last.push(item);
              return table;
            },
            [[]]
          )
          .map((dd, i) => {
            const rowClass = "row" + (i % 5 == 0 ? " row-pagebreak" : "");
            return (
              <div className={rowClass}>
                {dd.map((d, i) => {
                  const blob = new Blob([d.image], { type: "image/png" });
                  return (
                    <div className="card">
                      <div className="card-content">
                        <img
                          className="qrcode"
                          width="100px"
                          src={"data:image/png;base64," + d.image}
                        />
                        <div>{d.id}</div>
                        <div>{title}</div>
                      </div>
                    </div>
                  );
                })}
              </div>
            );
          })}
      </div>
    </div>
  );
};

export default App;
