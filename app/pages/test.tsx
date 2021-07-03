import React, { useState } from "react";

type DataStruct = {
  Name: string;
  EMail: string;
  Password: string;
};
const userData: DataStruct = {
  Name: "",
  EMail: "",
  Password: "",
};

const View = (): JSX.Element => {
  const ApiFetch = () => {
    fetch("http://localhost:8080/registrantion", {
      mode: "cors",
      method: "POST",
      headers: {"Content-Type": "application/json",}, // JSON形式のデータのヘッダー
      body: JSON.stringify(data),
    })
      .then((res) => res.json())
      .then((data) => {
        setPosts(data);
      });
  };
  const [data, setData] = useState<DataStruct>(userData);
  const [posts, setPosts] = useState([]);
  const onChangeName = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setData({ ...data, Name: value });
}
  const onChangeEMail = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value;
    setData({ ...data, EMail: value });
  }
  const onChangePassword = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value;
    setData({ ...data, Password: value });
  }
  return (
    <>
      <div>ユーザー登録して一緒に記事を投稿しましょう！</div>
      <label>名前</label>
      <input type="string" value={data.Name} onChange={onChangeName}></input>
      <label>EMail</label>
      <input type="string" value={data.EMail} onChange={onChangeEMail}></input>
      <label>Password</label>
      <input type="string" value={data.Password} onChange={onChangePassword}></input>
      <button onClick={ApiFetch}>送信</button>
      {JSON.stringify(posts)}
    </>
  );
};

export default View;