import {
  HStack,
  Stack,
  Button,
  Box,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
} from '@chakra-ui/react';
import React, { useState } from "react"
import { useRouter } from 'next/router'
import NextLink from "next/link";
import Link from "../components/Link"
import { useEffect } from 'react';
import Template from '../template';
type user = {
	ID: string
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Posts: string
	Profile: string
	ProfileID: string
	Goods: string
};
var returnData: user = {
	ID: "",
	CreatedAt: "",
	UpdatedAt: "",
	DeletedAt: "",
	Name: "",
	EMail: "",
	Password: "",
	Posts: "",
	Profile: "",
	ProfileID: "",
	Goods: "",
}
const Fuga = () => {
  const router = useRouter();
  const [answer, setAnswer] = useState<string>()
  const [choicesData, setChoicesData] = useState<any>()
  const [post, setpost] = useState()
  const [user, setUser] = useState<user>(returnData);
  console.log(user)
  useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.json())
			.then((data) => {
				const result = JSON.stringify(data)
				const result2: user = JSON.parse(result)
				setUser(result2)
			})
    },[])
  const DeleteFetch = () => {
    const ArticleData = { ID: Number(router.query.ID), UserID: Number(router.query.UserID) }
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
        {router.query.Name==user.Name && 
        <Menu>
          <MenuButton as={Button} h={10} p={2}>
            ...
          </MenuButton>
          <MenuList>
            <MenuItem>
            <NextLink
							as={`/post/${router.query.ID}`}
							href={{ pathname: `/post/[ID]`, query: router.query }}
							passHref>
							<Box>記事を編集する</Box>
						</NextLink></MenuItem>
            <MenuItem><Button onClick={DeleteFetch}>記事を削除する</Button></MenuItem>
          </MenuList>
        </Menu>}
        
        
        {post && <>{JSON.stringify(post)}</>}
      </Stack>
    </>)
}
export default Fuga