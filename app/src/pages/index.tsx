import {
	chakra,
	Stack,
	Box,
	VStack,
	Button,
	Image,
} from '@chakra-ui/react';
import NextLink from "next/link";
import Link from "./components/Link"
import React, { useState, useEffect } from "react"
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
const Home = (): JSX.Element => {
	const [postDatas, setPostDatas] = useState<post[]>([])
	useEffect(() => {
		fetch("http://localhost:8080/getalluserpost", {
			mode: "cors",
			method: "GET",
		}).then((res) => res.json())
			.then((datas) => {
				const userDatas: user[] = datas
				console.log("貰ってきたデータ", datas)
				userDatas.forEach(userData =>//user型データがいくつか入ってる
				{
					let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
					userData.Posts.forEach(userDataPost =>//user型データの中のpost
					{
						userDataPost.Name = userData.Name,
						userDataPost.Icon = blob
							setPostDatas(postDatas => [...postDatas, userDataPost])
					})
				})
			})
			.catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
	}, [])
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
						<Box bgColor="aquamarine"><Link href={`/${postData.Name}`}>{postData.Name}</Link>が{postData.CreatedAt.substring(0, 10)}に投稿しました</Box>
						<NextLink
							as={`/items/${userInfo.ID}`}
							href={{ pathname: `/items/[ID]`, query: userInfo }}
							passHref>
							<Box as="a" fontWeight="bold" bgColor="azure">{userInfo.Title}</Box>
						</NextLink>
					</VStack>
				)
			})}

		</>)
	}
	return (
		<>
			<chakra.div>
				<Template />
				クイズを投稿して知見を共有しませんか！
				<Stack direction="row" align="center">
					<Link href="/post">
						<Button colorScheme="teal" variant="solid">
							クイズを投稿する
						</Button>
					</Link>
				</Stack>
				<Box align="center" p="10">皆が投稿したクイズ一覧</Box>
				<Stack>
					<VStack>
						<PostsView />
					</VStack>
				</Stack>
			</chakra.div>
		</>
	)
}

export default Home