import { useState } from "react";
import * as React from "react";
import { useForm } from "react-hook-form";
import Link from './components/Link';
import {
  chakra,
  Input,
  FormLabel,
  Button,
  FormControl,
  FormHelperText,
  Checkbox,
  Box,
} from "@chakra-ui/react"
import MyPagesTemplate from "./myPagesTemplate";

type FormData = {
  Name: string;
  EMail: string;
  Password: string;
};
type FormType = {
  Name: string;
  EMail: string;
  Password: string;
  isTerms: boolean;
}
const userData: FormData = {
  Name: "",
  EMail: "",
  Password: "",
};
const View = () => {
  const { register, handleSubmit, formState, formState: { errors }, getValues } = useForm<FormType>({
    mode: "onTouched"
  });
  const [posts, setPosts] = useState([]);
  const setData = () => {
    const hasData = getValues(["Name", "EMail", "Password"]);
    userData.Name = hasData[0]
    userData.EMail = hasData[1]
    userData.Password = hasData[2]
    console.log(userData)
  };
  const ApiFetch = () => {
    setData()
    fetch("http://localhost:8080/registrantion", {
      mode: "cors",
      method: "POST",
      headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
      body: JSON.stringify(userData),
    })
      .then((res) => res.json())
      .then((data) => {
        setPosts(data);
      });
  };

  return (
    <>
      <chakra.div>
        <MyPagesTemplate />
        ユーザー登録して一緒に記事を投稿しましょう！
        <FormControl onSubmit={handleSubmit(setData)}
          isInvalid={errors.Name ? true : false}>
          <FormLabel>名前</FormLabel>
          <Input
            type="string"
            placeholder="田中　太郎"
            {...register("Name", {
              required: true,
              minLength: {
                value: 0,
                message: '名前を入力してください' // JS only: <p>error message</p> TS only support string
              }
            })}
          />
          {errors.Name && "名前を入力してください"}
        </FormControl>

        <FormControl onSubmit={handleSubmit(setData)}
          isInvalid={errors.EMail ? true : false}>
          <FormLabel>Email address</FormLabel>
          <Input
            type="email"
            placeholder="example@gmail.com"
            {...register("EMail", {
              required: "EMailを入力してください",
              pattern: {
                value: /^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/,
                message: 'メールアドレスを入力してください' // JS only: <p>error message</p> TS only support string
              },
            })}
          />
          {errors.EMail && errors.EMail.message}
          <FormHelperText>We will never share your email.</FormHelperText>
        </FormControl>

        <FormControl onSubmit={handleSubmit(setData)}
          isInvalid={errors.Password ? true : false}>
          <FormLabel>Password</FormLabel>
          <Input
            type="password"
            placeholder="password1"
            {...register("Password", {
              required: "パスワードを入力してください",
              minLength: {
                value: 8,
                message: '8文字以上にしてください' // JS only: <p>error message</p> TS only support string
              },
              pattern: {
                value: /^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d)[a-zA-Z\d]{8,100}$/,
                message: '小文字大文字数字をそれぞれ含めてください' // JS only: <p>error message</p> TS only support string
              }
            })}
          />
          {errors.Password && errors.Password.message}
          <FormHelperText>Passwordは8文字以上かつ小文字大文字数字を含めてください。We will never share your password.</FormHelperText>
        </FormControl>

        <FormControl onSubmit={handleSubmit(setData)}
          isInvalid={errors.isTerms ? true : false}>
          <Checkbox
            {...register("isTerms", { required: true })}
          ><Link color="teal.500" href="/terms">
              利用規約</Link>に同意する</Checkbox>
        </FormControl>
        <Button type="submit"
          colorScheme="teal"
          onClick={ApiFetch}
          disabled={!formState.isValid}
        >送信</Button>
        {JSON.stringify(posts)}
      </chakra.div>
    </>
  );
};

export default View;