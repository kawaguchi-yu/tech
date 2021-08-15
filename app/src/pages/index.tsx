import {
	chakra,
	Stack,
	Box,
	VStack,
	Image,
	Grid,
	GridItem,
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
	Goods: number
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
	Goods: good[]
	Icon: Blob
};
type good = {
	ID: number
	UserId: number
	PostID: number
}
const Home = (): JSX.Element => {
	const [postDatas, setPostDatas] = useState<post[]>([])
	useEffect(() => {
		fetch("http://localhost:8080/getalluserpost", {
			mode: "cors",
			method: "GET",
		}).then((res) => res.json())
			.then((datas) => {
				var returnPostDatas: post[]
				const userDatas: user[] = datas
				console.log("貰ってきたデータ", datas)
				userDatas.forEach(userData =>{
					let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
					userData.Posts.forEach(postData =>{
						postData.Name = userData.Name,
						postData.Icon = blob
						if (returnPostDatas == null) {
							returnPostDatas = [postData]
						} else {
							returnPostDatas = [...returnPostDatas, postData]
						}
					})
				})
				setPostDatas(returnPostDatas)
			})
			.catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
	}, [])
	const PostsView = () => {
		return (<>
			{postDatas.map((postData) => {
				const userInfo = {
					Title: postData.Title,
				}
				return (
					<VStack key={postData.ID} padding="10" bg="white" boxShadow="xs">
						<Image boxSize="50px"
							borderRadius="full"
							src={(window.URL.createObjectURL(postData.Icon))}
							alt="select picture" />
						<Box bgColor="aquamarine"><Link href={`/${postData.Name}`}>{postData.Name}</Link>が{postData.CreatedAt.substring(0, 10)}に投稿しました</Box>
						<NextLink
							as={`/items/${postData.ID}`}
							href={{ pathname: `/items/[ID]`, query: userInfo }}
							passHref>
							<Box as="a" fontWeight="bold" bgColor="azure">{userInfo.Title}</Box>
						</NextLink>
						<Box>いいね数{postData.Goods?postData.Goods.length:0}</Box>
					</VStack>
				)
			})}

		</>)
	}
	return (
		<>
			<chakra.div>
				<Template />
				<Stack>
				<Grid
					h="200px"
					templateRows="repeat(1, 1fr)"
					templateColumns="repeat(5, 1fr)"
					margin={5}
					gap={2}
				>
					<GridItem rowSpan={2} colSpan={1}>いいね順に並び変える</GridItem>
					<GridItem rowSpan={2} colSpan={4} align="center">皆が投稿したクイズ一覧<PostsView /></GridItem>
					</Grid>
				</Stack>
			</chakra.div>
		</>
	)
}

export default Home