import React, { useState } from "react"
import { useForm } from "react-hook-form";
import {
	Box,
	FormControl,
	FormLabel,
	Input,
	Checkbox,
	HStack,
	Stack,
	Button,
	RadioGroup,
	Radio,
} from '@chakra-ui/react';
import Link from './components/Link';
import Template from "./template";
type quizType = {
	title: string
	answer: string
	wrongAnswer1: string
	wrongAnswer2: string
	wrongAnswer3: string
}
const quizData: quizType = {
	title: "",
	answer: "",
	wrongAnswer1: "",
	wrongAnswer2: "",
	wrongAnswer3: "",
}
const Posts = () => {
	const { register, handleSubmit, formState, formState: { errors }, getValues } = useForm({
		mode: "onTouched",
	});
	const [posts, setPosts] = useState([])

	const setData = () => {
		const hasData = getValues(["title", "text1", "text2", "text3", "text4"])
		quizData.title = hasData[0]
		quizData.answer = hasData[1]
		quizData.wrongAnswer1 = hasData[2]
		quizData.wrongAnswer2 = hasData[3]
		quizData.wrongAnswer3 = hasData[4]
	}
	const test = () => {
		setData()
		console.log(quizData)
	}
	const ApiFetch = () => {
		setData()
		fetch("http://localhost:8080/post", {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: JSON.stringify(quizData),
		})
			.then((res) => res.json())
			.then((data) => {
				setPosts(data);
			})
			.catch((err) => { console.log(err) })
	};
	return (<>
		<Template />
		<FormControl onSubmit={handleSubmit(setData)}
			isInvalid={errors.title ? true : false}>
			<FormLabel>タイトル</FormLabel>
			<Input
				type="string"
				placeholder="例:フロントエンド言語は？"
				{...register("title", {
					required: true,
					minLength: {
						value: 0,
						message: '例を入力してください' // JS only: <p>error message</p> TS only support string
					}
				})}
			/>
			{errors.title && "タイトルを入力してください"}
		</FormControl>
		<HStack>
			<FormControl onSubmit={handleSubmit(setData)}
				isInvalid={errors.answer ? true : false}>
				<FormLabel>正答</FormLabel>
				<Input
					type="body"
					placeholder="例:JavaScript"
					{...register("answer", {
						required: "回答を入力してください",
					})}
				/>
				{errors.answer && errors.answer.message}
			</FormControl>

			<FormControl onSubmit={handleSubmit(setData)}
				isInvalid={errors.wrongAnswer1 ? true : false}>
				<FormLabel>誤答</FormLabel>
				<Input
					type="body"
					placeholder="例:Java"
					{...register("wrongAnswer1", {
						required: "回答を入力してください",
					})}
				/>
				{errors.wrongAnswer1 && errors.wrongAnswer1.message}
			</FormControl>

			<FormControl onSubmit={handleSubmit(setData)}
				isInvalid={errors.wrongAnswer2 ? true : false}>
				<FormLabel>誤答</FormLabel>
				<Input
					type="body"
					placeholder="例:PHP"
					{...register("wrongAnswer2", {
						required: "回答を入力してください",
					})}
				/>
				{errors.wrongAnswer2 && errors.wrongAnswer2.message}
			</FormControl>

			<FormControl onSubmit={handleSubmit(setData)}
				isInvalid={errors.wrongAnswer3 ? true : false}>
				<FormLabel>誤答</FormLabel>
				<Input
					type="body"
					placeholder="例:Ruby"
					{...register("wrongAnswer3", {
						required: "回答を入力してください",
					})}
				/>
				{errors.wrongAnswer3 && errors.wrongAnswer3.message}
			</FormControl>
		</HStack>
		<Button type="submit"
			colorScheme="teal"
			onClick={test}
			disabled={!formState.isValid}
		>送信</Button>
		{JSON.stringify(posts)}
	</>)
}
export default Posts