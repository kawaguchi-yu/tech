import {
  Stack,
  Button,
  Box,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  VStack,
  Heading,
  Container,
  GridItem,
  useColorModeValue,
} from '@chakra-ui/react';
import React, { useState, useEffect } from "react"
import { useRouter } from 'next/router'
import NextLink from "next/link";
import Link from "../../../public"
import { StarIcon } from '@chakra-ui/icons'
import Template from '../template';
import { sessionInformation } from '../../../env'
type user = {
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string
  Name: string
  EMail: string
  Password: string
  Posts: post[]
  Profile: string
  ProfileID: string
};
type post = {
  Name: string//userの名前を入れる
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string
  UserID: number
  Title: string
  Answer: string
  WrongAnswer1: string
  WrongAnswer2: string
  WrongAnswer3: string
  Explanation: string
  Goods: good[]
  Icon: Blob
};
type good = {
  ID: number
  UserID: number
  PostID: number
}
var returnData: user = {
  ID: 0,
  CreatedAt: "",
  UpdatedAt: "",
  DeletedAt: "",
  Name: "",
  EMail: "",
  Password: "",
  Posts: [{
    Name: "",
    ID: 0,
    CreatedAt: "",
    UpdatedAt: "",
    DeletedAt: "",
    UserID: 0,
    Title: "",
    Answer: "",
    WrongAnswer1: "",
    WrongAnswer2: "",
    WrongAnswer3: "",
    Explanation: "",
    Goods: [{ ID: 0, UserID: 0, PostID: 0 }],
    Icon: null,
  }],
  Profile: "",
  ProfileID: "",
}
type URLPath = {
  UserID: number
}
const Fuga = () => {
  const router = useRouter();
  const [URLQuery, setURLQuery] = useState<URLPath>()
  const [answer, setAnswer] = useState<string>()
  const [choicesData, setChoicesData] = useState<string[]>()
  const [userInPost, setUserInPost] = useState<user>(returnData)
  const [user, setUser] = useState<user>(returnData);
  const [isGooded, setIsGooded] = useState<boolean>(false)
  const [uiGoodedCount, setUiGoodedCount] = useState<number>(0)
  useEffect(() => {
    if (router.asPath !== router.route) {//厳密不等価
      const queryID: URLPath = { UserID: Number(router.query.ID) }
      setURLQuery(queryID);
    }
  }, [router])
  useEffect(() => {
    fetch(`${sessionInformation.backendHost}/user`, {
      mode: "cors",
      method: "GET",
      credentials: 'include',
    }).then((res) => res.json())
      .then((data) => {
        const result = JSON.stringify(data)
        const result2: user = JSON.parse(result)
        setUser(result2)
      })
  }, [])
  useEffect(() => {
    if (URLQuery) {
      fetch(`${sessionInformation.backendHost}/getuserbyid`, {
        mode: "cors",
        method: "POST",
        headers: { "Content-Type": "application/json", },
        credentials: 'include',
        body: JSON.stringify(URLQuery)
      }).then((res) => res.json())
        .then((data) => {
          const result = JSON.stringify(data)
          const result2: user = JSON.parse(result)
          setUserInPost(result2)
          console.log(result2)
        })
    }
  }, [URLQuery])
  const DeletePostFetch = () => {
    const ArticleData = { ID: Number(router.query.ID), UserID: Number(user.ID) }
    fetch(`${sessionInformation.backendHost}/deletepost`, {
      mode: "cors",
      method: "POST",
      headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
      credentials: 'include',
      body: JSON.stringify(ArticleData),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
      })
      .catch((err) => { console.log(err) })
  };
  const DeleteGoodFetch = () => {
    const GoodData = { userID: Number(user.ID), postID: Number(router.query.ID) }
    fetch(`${sessionInformation.backendHost}/deletegood`, {
      mode: "cors",
      method: "POST",
      headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
      credentials: 'include',
      body: JSON.stringify(GoodData),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
      })
      .catch((err) => { console.log(err) })
  }
  const GoodFetch = () => {

    const GoodData = { userID: Number(user.ID), postID: Number(router.query.ID) }
    console.log("GoodData", GoodData)
    fetch(`${sessionInformation.backendHost}/good`, {
      mode: "cors",
      method: "POST",
      headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
      credentials: 'include',
      body: JSON.stringify(GoodData),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
      })
      .catch((err) => { console.log(err) })
  }
  const RandomAnswer = () => {
    const choices = [
      userInPost.Posts[0].Answer,
      userInPost.Posts[0].WrongAnswer1,
      userInPost.Posts[0].WrongAnswer2,
      userInPost.Posts[0].WrongAnswer3,
    ]
    for (let i = choices.length - 1; i > 0; i--) {
      let j = Math.floor(Math.random() * (i + 1));
      let tmp = choices[i];
      choices[i] = choices[j];
      choices[j] = tmp;
    }
    setChoicesData(choices)
  }
  const GoodCheck = () => {
    if (userInPost.Posts[0].Goods) {
      userInPost.Posts[0].Goods.map(goodData => {
        console.log("a", goodData.UserID, "=", user.ID)
        if (goodData.UserID == user.ID && user.ID != 0) {
          setIsGooded(true)
          console.log("isgoodedがtrueになりました", goodData.UserID, "=", user.ID)
        }
      })
    }
  }
  useEffect(() => { RandomAnswer(), GoodCheck() }, [userInPost])
  const getAnswer = (event) => {
    if (userInPost.Posts[0].Answer == event.target.value) {
      setAnswer("正解！")
    } else {
      setAnswer("不正解！")
    }
  }
  return (
    <>
      <Template />
      <Stack align="center">
        <VStack
          bg={useColorModeValue('white', 'gray.700')}
          boxShadow={'lg'}
          p={8}>
          <Heading>{userInPost.Posts[0].Title}</Heading>
          {choicesData && <><Container>回答1:{choicesData[0]}</Container><Button margin="2" onClick={getAnswer} value={choicesData[0]}>これにする</Button></>}
          {choicesData && <><Container>回答1:{choicesData[1]}</Container><Button margin="2" onClick={getAnswer} value={choicesData[1]}>これにする</Button></>}
          {choicesData && <><Container>回答1:{choicesData[2]}</Container><Button margin="2" onClick={getAnswer} value={choicesData[2]}>これにする</Button></>}
          {choicesData && <><Container>回答1:{choicesData[3]}</Container><Button margin="2" onClick={getAnswer} value={choicesData[3]}>これにする</Button></>}
        </VStack>
        <Stack>
          <>{answer}</>
          {answer == "不正解！" && <>正解は{userInPost.Posts[0].Answer}です</>}
        </Stack>
        <Stack>
          {answer && <>解説文:  {userInPost.Posts[0].Explanation}</>}
        </Stack>

        <Link href="/">
          <Box>戻る</Box>
        </Link>
        {userInPost.Name == user.Name &&
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
                  <Button>記事を編集する</Button>
                </NextLink></MenuItem>
              <MenuItem><Button onClick={DeletePostFetch}>記事を削除する</Button></MenuItem>
            </MenuList>
          </Menu>}
        {userInPost.Posts[0].UserID == user.ID
          ? <Button disabled><StarIcon />いいね数{userInPost.Posts[0].Goods ? userInPost.Posts[0].Goods.length + uiGoodedCount : uiGoodedCount}</Button>
          : isGooded == true
            ? <Button onClick={() => { DeleteGoodFetch(), setIsGooded(false), setUiGoodedCount(uiGoodedCount - 1) }}><StarIcon color="gold" />いいねしました。{userInPost.Posts[0].Goods ? userInPost.Posts[0].Goods.length + uiGoodedCount : uiGoodedCount}</Button>
            : <><Button onClick={() => { GoodFetch(), setIsGooded(true), setUiGoodedCount(uiGoodedCount + 1) }} disabled={!user.ID} ><StarIcon />いいねする{userInPost.Posts[0].Goods ? userInPost.Posts[0].Goods.length + uiGoodedCount : uiGoodedCount}</Button>
              {!user.ID && <Box>ログインしないといいねができません</Box>}
            </>}
      </Stack>
    </>)
}
export default Fuga