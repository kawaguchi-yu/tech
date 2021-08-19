import React, { useState, useEffect } from "react"
import { useForm } from "react-hook-form";
import {
	FormControl,
	FormLabel,
	useToast,
	Input,
	VStack,
	HStack,
	Stack,
	Button,
	Container,
	Box,
} from '@chakra-ui/react';
import Template from "../template";
import { useRouter } from 'next/router'
type user = {
	ID: number
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Posts: post
	Profile: string
	ProfileID: string
};
type post = {
	ID: number
	UserID: number
	Title: string
	Answer: string
	WrongAnswer1: string
	WrongAnswer2: string
	WrongAnswer3: string
	Explanation: string
}
const quizData: post = {
	ID: 0,
	UserID: 0,
	Title: "",
	Answer: "",
	WrongAnswer1: "",
	WrongAnswer2: "",
	WrongAnswer3: "",
	Explanation: "",
}
type URLPath = {
	UserID: number
}
const Posts = () => {
	const { register, handleSubmit, formState, formState: { errors }, setValue, getValues } = useForm({
		mode: "onTouched",
	});
	const [URLQuery, setURLQuery] = useState<URLPath>()
	const [randomAnswer, setRandomAnswer] = useState([])
	const [answer, setAnswer] = useState<string>()
	const [getPost, setGetpost] = useState<post>(quizData)
	const router = useRouter();
	const toast = useToast()
	useEffect(() => {
		if (router.asPath !== router.route) {//厳密不等価
			const queryID: URLPath = { UserID: Number(router.query.ID) }
			setURLQuery(queryID);
		}
	}, [router])
	useEffect(() => {
		if (URLQuery) {
			console.log("URLQuery", URLQuery)
			fetch("http://localhost:8080/getuserbyid", {
				mode: "cors",
				method: "POST",
				headers: { "Content-Type": "application/json", },
				credentials: 'include',
				body: JSON.stringify(URLQuery)
			}).then((res) => res.json())
				.then((data) => {
					const result = JSON.stringify(data)
					const result2: user = JSON.parse(result)
					setGetpost(result2.Posts[0])
					setFetchData(result2.Posts[0])
					console.log("result", result2.Posts[0])
				})
		}
	}, [URLQuery])

	const upDateFetch = () => {
		setData()
		fetch("http://localhost:8080/updatepost", {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: JSON.stringify(quizData),
		})
			.then((res) => res.json())
			.then((data) => {
				console.log(data);
				if(data!="正常に終了しました"){
					toast({
						title: "エラーが発生しました",
						description: "編集権限がありません",
						status: "error",
						duration: 9000,
						isClosable: true,
					})
					return
				}
				toast({
					title: "記事を更新しました！",
					description: "正常に記事が更新されました。",
					status: "success",
					duration: 9000,
					isClosable: true,
				})
			})
			.catch((err) => { console.log(err) })
	};
	const setFetchData = (postData) => {
		setValue("title", postData.Title)
		setValue("answer", postData.Answer)
		setValue("wrongAnswer1", postData.WrongAnswer1)
		setValue("wrongAnswer2", postData.WrongAnswer2)
		setValue("wrongAnswer3", postData.WrongAnswer3)
		setValue("explanation", postData.Explanation)
	}
	const setData = () => {
		const hasData = getValues(["title", "answer", "wrongAnswer1", "wrongAnswer2", "wrongAnswer3", "explanation"])
		quizData.ID = getPost.ID
		quizData.UserID = getPost.UserID
		quizData.Title = hasData[0]
		quizData.Answer = hasData[1]
		quizData.WrongAnswer1 = hasData[2]
		quizData.WrongAnswer2 = hasData[3]
		quizData.WrongAnswer3 = hasData[4]
		quizData.Explanation = hasData[5]
		console.log(quizData)
	}
	const RandomAnswer = () => {
		let answer = [quizData.Answer, quizData.WrongAnswer1, quizData.WrongAnswer2, quizData.WrongAnswer3]
		for (let i = answer.length - 1; i > 0; i--) {
			let j = Math.floor(Math.random() * (i + 1));
			let tmp = answer[i];
			answer[i] = answer[j];
			answer[j] = tmp;
		}
		setRandomAnswer(answer)
	}
	const test = () => {
		setData()
		RandomAnswer()
	}
	const getAnswer = (event) => {
		if (quizData.Answer == event.target.value) {
			setAnswer("正解！")
		} else {
			setAnswer("不正解！")
		}
	}

	return (<>
		<Template />
		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.title ? true : false}>
			<FormLabel>問題文</FormLabel>
			<Input
				type="string"
				defaultValue={getPost.Title}
				placeholder="例:この中でフロントエンド言語はどれ？"
				{...register("title", {
					required: true,
					pattern: {
						value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
						message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.title && errors.title.message}
		</FormControl>
		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.answer ? true : false}>
			<FormLabel>正答</FormLabel>
			<Input
				type="body"
				defaultValue={getPost.Answer}
				placeholder="例:JavaScript"
				{...register("answer", {
					required: true,
					pattern: {
						value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
						message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.answer && errors.answer.message}
		</FormControl>
		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.wrongAnswer1 ? true : false}>
			<FormLabel>誤答</FormLabel>
			<Input
				type="body"
				defaultValue={getPost.WrongAnswer1}
				placeholder="例:Go"
				{...register("wrongAnswer1", {
					required: true,
					pattern: {
						value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
						message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.wrongAnswer1 && errors.wrongAnswer1.message}
		</FormControl>

		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.wrongAnswer2 ? true : false}>
			<FormLabel>誤答</FormLabel>
			<Input
				type="body"
				defaultValue={getPost.WrongAnswer2}
				placeholder="例:PHP"
				{...register("wrongAnswer2", {
					required: true,
					pattern: {
						value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
						message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.wrongAnswer2 && errors.wrongAnswer2.message}
		</FormControl>

		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.wrongAnswer3 ? true : false}>
			<FormLabel>誤答</FormLabel>
			<Input
				type="body"
				defaultValue={getPost.WrongAnswer3}
				placeholder="例:Ruby"
				{...register("wrongAnswer3", {
					required: true,
					pattern: {
						value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
						message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.wrongAnswer3 && errors.wrongAnswer3.message}
		</FormControl>
		<Stack>
			<FormControl onSubmit={handleSubmit(setData)}
				isInvalid={errors.explanation ? true : false}>
				<FormLabel>解説文</FormLabel>
				<Input
					type="body"
					defaultValue={getPost.Explanation}
					placeholder="JavaScriptだけがフロントエンド言語だよ！"
					{...register("explanation", {
						required: true,
						pattern: {
							value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,200}$/,
							message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
						}
					})}
				/>
				{errors.explanation && errors.explanation.message}
			</FormControl>
		</Stack>
		<Button type="submit"
			colorScheme="teal"
			onClick={upDateFetch}
			disabled={!formState.isValid}
		>更新</Button>
			<Button onClick={test}>プレビュー</Button>
		<Stack>
			<>問題文:{quizData.Title}</>
		</Stack>
		<VStack>
		<Container>回答1:{randomAnswer[0]}<Button margin="2" onClick={getAnswer} value={randomAnswer[0]}>これにする</Button></Container>
		<Container>回答2:{randomAnswer[1]}<Button margin="2" onClick={getAnswer} value={randomAnswer[1]}>これにする</Button></Container>
		<Container>回答3:{randomAnswer[2]}<Button margin="2" onClick={getAnswer} value={randomAnswer[2]}>これにする</Button></Container>
		<Container>回答4:{randomAnswer[3]}<Button margin="2" onClick={getAnswer} value={randomAnswer[3]}>これにする</Button></Container>
		</VStack>
		<Stack>
			<>{answer}</>
			{answer && <>正解は<Box>{quizData.Answer}です</Box></>}
		</Stack>
		<Stack>
			{answer && <>解説<Box> {quizData.Explanation}</Box></>}
		</Stack>
	</>)
}
export default Posts