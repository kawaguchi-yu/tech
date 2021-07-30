import React, { useState, useEffect } from "react"
import { useRouter } from "next/router"
import {
	chakra,
	Stack,
	Box,
	VStack,
} from '@chakra-ui/react';
import Link from "next/link";
import Template from "./template";
type user = {
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
	Tags: any
	Goods: any
};
type URLPath = {
	Name: string | string[]
}
const MyPages = (): JSX.Element => {

	const router = useRouter()
	const [URLQuery, setURLQuery] = useState<URLPath>()
	useEffect(() => {
		if (router.asPath !== router.route) {
			const queryname: URLPath = { Name: router.query.name }
			setURLQuery(queryname);
		}
	}, [router])
	useEffect(() => {
		if (URLQuery) {
			fetch("http://localhost:8080/getuserpost", {
				mode: "cors",
				method: "POST",
				headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
				credentials: 'include',//bodyの代わりにcookieを送る
				body: JSON.stringify(URLQuery)
			})
			.then((res) => res.json())
			.then((data) => {
				data.forEach(array =>
					setPostDatas(postDatas => [...postDatas, array]),
					console.log("貰ってきたデータ", data)
				)
			})
			.catch(() => {
				console.error("データを貰ってくることができませんでした")
				console.log("URLQuery", URLQuery)
			})
		}
	}, [URLQuery])
	const [postDatas, setPostDatas] = useState<user[]>([])
	const PostsView = () => {
		return (<>
			{postDatas.map((postData) => {
				const userInfo: user = {
					ID: postData.ID,
					CreatedAt: postData.CreatedAt,
					UpdatedAt: postData.UpdatedAt,
					DeletedAt: postData.DeletedAt,
					UserID: postData.UserID,
					Title: postData.Title,
					Answer: postData.Answer,
					WrongAnswer1: postData.WrongAnswer1,
					WrongAnswer2: postData.WrongAnswer2,
					WrongAnswer3: postData.WrongAnswer3,
					Explanation: postData.Explanation,
					Tags: postData.Tags,
					Goods: postData.Goods,
				}
				return (
					<VStack key={userInfo.ID} padding="10" border="solid 1px">
						<Box>{postData.CreatedAt.substring(0, 10)}</Box>
						<Link
							as={`/items/${userInfo.ID}`}
							href={{ pathname: `/items/[ID]`, query: userInfo }}
							passHref>
							<Box>{userInfo.Title}</Box>
						</Link>
					</VStack>
				)
			})}

		</>)

	}
	return (
		<>
			<Template />
			<chakra.div>
				<Box align="center" p="10">投稿したクイズ一覧</Box>
				<Stack>
					<VStack>
						<PostsView />
					</VStack>
				</Stack>
				<>{router.query.name}</>
			</chakra.div>
		</>
	)
}
export default MyPages