import React, { useState, useEffect } from "react"
import { useRouter } from "next/router"
import {
	chakra,
	Stack,
	Box,
	VStack,
	HStack,
	Image,
} from '@chakra-ui/react';
import Link from "next/link";
import Template from "./template";
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
	Goods: good[]
	Icon: Blob
};
type good ={
	ID:number
	UserId:number
	PostID:number
}
type URLPath = {
	Name: string
}
const MyPages = (): JSX.Element => {
	const router = useRouter()
	const [URLQuery, setURLQuery] = useState<URLPath>()
	const [postDatas, setPostDatas] = useState<post[]>([])
	const [userId, setUserId] = useState<good>()
	const [goodedUserDatas,setGoodedUserDatas] = useState<user[]>([])
	const [goodedPostDatas,setGoodedPostDatas] = useState<post[]>([])
	useEffect(() => {
		if (router.asPath !== router.route) {//厳密不等価
			const queryname: URLPath = { Name: String(router.query.name) }
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
			}).then((res) => res.json())
				.then((datas) => {
					const userData: user = datas
					console.log("貰ってきたデータ",userData)
					let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
					userData.Posts.forEach(postData => {
						postData.Name = userData.Name
						postData.Icon = blob
					})
					console.log("貰ってきたデータ", datas)
					setPostDatas(userData.Posts)
					setUserId({ID:0,PostID:0,UserId:userData.ID})
				})
				.catch(() => {
					console.error("データを貰ってくることができませんでした")
				})
		}
	}, [URLQuery])
	useEffect(()=>{
		if (userId){
		fetch("http://localhost:8080/returngoodedpost", {
				mode: "cors",
				method: "POST",
				headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
				credentials: 'include',//bodyの代わりにcookieを送る
				body: JSON.stringify(userId)
			}).then((res) => res.json())
			.then((datas) => {
				var returnPostDatas:post[]
				console.log("貰ってきたデータ(aaaaaaaaa)",datas)
				const userDatas: user[] = datas
				userDatas.forEach(userData =>{
					let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
				userData.Posts.forEach(postData => {
					postData.Name = userData.Name
					postData.Icon = blob
					if(returnPostDatas == null){
						returnPostDatas = [postData]
					}else{
					returnPostDatas =[...returnPostDatas,postData]
				}})
				console.log("貰ってきたデータ222222222222", datas)
			})
			setGoodedPostDatas(returnPostDatas)
		})
			.catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
		}},[userId])
	const PostsView = () => {
		return (<>
			{postDatas.map((postData) => {
				return (
					<VStack key={postData.ID} padding="10" border="solid 1px">
						<Image boxSize="50px"
							borderRadius="full"
							src={(window.URL.createObjectURL(postData.Icon))}
							alt="select picture" />
							{postData.Goods ?<Box>いいね数:{postData.Goods.length}</Box>:<Box>いいね数:0</Box>}
						<Box>{postData.Name}が{postData.CreatedAt.substring(0, 10)}に投稿しました</Box>
						<Link
							as={`/items/${postData.ID}`}
							href={{ pathname: `/items/[ID]`}}
							passHref>
							<Box>{postData.Title}{postData.UserID}</Box>
						</Link>
					</VStack>
				)
			})}
		</>)
	}
	const GoodedView = () => {
		return (<>
		{goodedPostDatas.map((goodedPostData)=>{
			console.log("goodedpostdatas",goodedPostData)
			return (
				<VStack key={goodedPostData.ID} padding="10" border="solid 1px">
					<Image boxSize="50px"
							borderRadius="full"
							src={(window.URL.createObjectURL(goodedPostData.Icon))}
							alt="select picture" />
							{goodedPostData.Goods ?<Box>いいね数:{goodedPostData.Goods.length}</Box>:<Box>いいね数:0</Box>}
						<Box><Link href={`/${goodedPostData.Name}`}>{goodedPostData.Name}</Link>が{goodedPostData.CreatedAt.substring(0, 10)}に投稿しました</Box>
						<Link
							as={`/items/${goodedPostData.ID}`}
							href={{ pathname: `/items/[ID]`}}
							passHref>
							<Box>{goodedPostData.Title}</Box>
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
			<Link href="/config">プロフィールを編集する</Link>
				<HStack align="center" p="10">
					<Box>{URLQuery&&URLQuery.Name}の投稿したクイズ一覧</Box>
					<Box>{URLQuery&&URLQuery.Name}のいいねしたクイズ一覧</Box>
				</HStack>
				<HStack>
					<VStack>
						<PostsView />
					</VStack>
					<VStack>
						<GoodedView />
					</VStack>
				</HStack>
			</chakra.div>
		</>
	)
}
export default MyPages