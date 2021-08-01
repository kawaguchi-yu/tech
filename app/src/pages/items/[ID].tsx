import {
  HStack,
  Stack,
  Button,
  Box,
} from '@chakra-ui/react';
import React, { useState } from "react"
import { useRouter } from 'next/router'
import Link from "../components/Link"
import { useEffect } from 'react';
import Template from '../template';
const Fuga = () => {
  const router = useRouter();
  console.log(router)
  const [answer, setAnswer] = useState<string>()
  const [choicesData, setChoicesData] = useState<any>()
  const [post, setpost] = useState()
  const DeleteFetch = () => {
    const ArticleData = {ID: Number(router.query.ID), UserID: Number(router.query.UserID)}
    console.log(ArticleData)
    fetch("http://localhost:8080/deletepost", {
      mode: "cors",
      method: "POST",
      headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
      credentials: 'include',
      body: JSON.stringify(ArticleData),
    })
      .then((res) => res.json())
      .then((data) => {
        setpost(data)
      })
      .catch((err) => { console.log(err) })
  };
  const RandomAnswer = () => {
    const choices = [
      router.query.Answer,
      router.query.WrongAnswer1,
      router.query.WrongAnswer2,
      router.query.WrongAnswer3,
    ]
    for (let i = choices.length - 1; i > 0; i--) {
      let j = Math.floor(Math.random() * (i + 1));
      let tmp = choices[i];
      choices[i] = choices[j];
      choices[j] = tmp;
    }
    setChoicesData(choices)
  }
  useEffect(RandomAnswer, [])
  const getAnswer = (event) => {
    if (router.query.Answer == event.target.value) {
      setAnswer("正解！")
    } else {
      setAnswer("不正解！")
    }
  }
  return (
    <>
      <Box><Template /></Box>
      <Stack align="center">
        <Stack>
          <Box>記事のID　{router.query.ID}</Box>
          <Box>記事のUserID　{router.query.UserID}</Box>
          <Box>記事制作者　{router.query.Name}</Box>
          <Box>問題:　　{router.query.Title}</Box>
        </Stack>
        <HStack>
          {choicesData && <>
            <Button onClick={getAnswer} value={choicesData[0]}>回答1:{choicesData[0]}</Button>
            <Button onClick={getAnswer} value={choicesData[1]}>回答2:{choicesData[1]}</Button>
            <Button onClick={getAnswer} value={choicesData[2]}>回答3:{choicesData[2]}</Button>
            <Button onClick={getAnswer} value={choicesData[3]}>回答4:{choicesData[3]}</Button>
          </>}</HStack>
        <Stack>
          <>{answer}</>
          {answer == "不正解！" && <>正解は{router.query.Answer}です</>}
        </Stack>
        <Stack>
          {answer && <>解説文:  {router.query.Explanation}</>}
        </Stack>

        <Link href="/">
          <Box>戻る</Box>
        </Link>
        <Button onClick={DeleteFetch}>記事を削除する</Button>
        {post && <>{JSON.stringify(post)}</>}
      </Stack>
    </>)
}
export default Fuga