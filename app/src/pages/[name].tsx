import React, { useState, useEffect } from "react"
import { useRouter } from "next/router"
import {
	chakra,
	Stack,
	Box,
	VStack,
	Image,
} from '@chakra-ui/react';
import Link from "next/link";
import Template from "./template";
type user = {
	ID: string
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Posts: post[]
	Profile: string
	ProfileID: string
	Goods: string
	Icon: string
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
	Tags: any
	Goods: any
	Icon: Blob
};
type URLPath = {
	Name: string
}
const MyPages = (): JSX.Element => {

	const router = useRouter()
	const [URLQuery, setURLQuery] = useState<URLPath>()
	useEffect(() => {
		if (router.asPath !== router.route) {//厳密不等価
			const queryname: URLPath = { Name: String(router.query.name)}
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
			.then((datas) => {
				const userData: user = datas
				let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
				userData.Posts.forEach(postData =>
					{postData.Name = userData.Name
						postData.Icon = blob
					setPostDatas(postDatas => [...postDatas, postData]),
					console.log("貰ってきたデータ", datas)
					})
			})
			.catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
		}
	}, [URLQuery])
	const [postDatas, setPostDatas] = useState<post[]>([])
	const PostsView = () => {
		return (<>
			{postDatas.map((postData) => {
				const userInfo = {
					Name: postData.Name,
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
						<Image boxSize="50px"
							borderRadius="full"
							src={(window.URL.createObjectURL(postData.Icon))}
							alt="select picture" />
						<Box>{router.query.name}が{postData.CreatedAt.substring(0, 10)}に投稿しました</Box>
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
			</chakra.div>
		</>
	)
}
export default MyPages