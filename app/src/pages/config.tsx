import React, { useState, useEffect } from "react"
import {
	Box,
	Button,
	Input,
	Image,
	Stack,
	Spacer,
	FormControl,
	FormLabel,
	InputGroup,
	InputLeftAddon,
	FormHelperText,
} from '@chakra-ui/react';
import router from "next/router";
import Template from "./template";
import { useForm } from "react-hook-form";
type user = {
	ID: number
	Name: string
	EMail: string
	Password: string
	Profile: profile
	ProfileID: number
	Icon: string
	IconBlob: Blob
};
type profile = {
	ID: number
	UserID: number
	Essay: string
	URLs: URL[]
}
type URL = {
	Name: string
	URL: string
	ProfileID: number
}
const userForm: user = {
	ID: 0,
	Name: "",
	EMail: "",
	Password: "",
	Profile: { ID: 0, UserID: 0, Essay: "", URLs: [{ Name: "", URL: "", ProfileID: 0 }, { Name: "", URL: "", ProfileID: 0 }] },
	ProfileID: 0,
	Icon: "",
	IconBlob: null,
}
const guestuser: number = 30;
const Config = () => {
	const { register, handleSubmit, formState, formState: { errors }, getValues } = useForm({
		mode: "onTouched",
	});
	const setData = () => {
		const hasData = getValues(["name", "essay", "url1", "url2"])
		userForm.ID = user.ID
		userForm.Name = hasData[0]
		userForm.Profile.ID = user.Profile.ID
		userForm.Profile.UserID = user.ID
		userForm.Profile.Essay = hasData[1]
		userForm.Profile.URLs[0].Name = "twitter"
		userForm.Profile.URLs[0].URL = hasData[2]
		userForm.Profile.URLs[1].Name = "github"
		userForm.Profile.URLs[1].URL = hasData[3]
		console.log(userForm)
	}
	const [view, setview] = useState<string>();
	const [iconData, setIconData] = useState<FormData>();
	const [posts, setPosts] = useState<Blob>();
	const [user, setUser] = useState<user>(userForm);
	useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.json())
			.then((data) => {
				const userData: user = data
				let bin = atob(userData.Icon.replace(/^.*,/, ''));
				let buffer = new Uint8Array(bin.length);
				for (let i = 0; i < bin.length; i++) {
					buffer[i] = bin.charCodeAt(i);
				} let blob = new Blob([buffer.buffer], {
					type: "image/jpeg"
				});
				userData.IconBlob = blob
				console.log(userData)
				setUser(userData)
			}).catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
	}, [])
	const onFileInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setview(window.URL.createObjectURL(e.target.files[0]))
		const image = new FormData()
		image.append("file", e.target.files[0])
		setIconData(image)
		console.log("ターゲットファイルの中身", e.target.files[0])
	};
	const ApiFetch = () => {
		const options: RequestInit = {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "multipart/form-data", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: iconData,
		}
		delete options.headers["Content-Type"];
		fetch("http://localhost:8080/seticon", options)
			.then((res) => res.blob())
			.then((data) => {
				setPosts(data);
				console.log("返ってきたデータ", data)
			})
			.catch((err) => { console.log(err) })
		console.log("アイコンデータ", iconData)
	};
	const updateUser = () => {
		setData()
		fetch("http://localhost:8080/updateuser", {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: JSON.stringify(userForm),
		}).then((res) => res.json())
			.then((data) => {
				console.log("送られてきたデータ" + data)
			})
			.catch((err) => { console.log(err) })
	}
	const DeleteFetch = () => {
		fetch("http://localhost:8080/deleteuser", {
			mode: "cors",
			method: "GET",
			credentials: "include",
		}).then((res) => res.json())
			.then((data) => {
				console.log(data)
			})
			.catch((err) => { console.log(err) })
		router.push("/")
	}
	return (<>
		<Template />
		{user.ID == guestuser && <Box bgColor="aquamarine">ゲストユーザーはアカウントを削除することができません。</Box>}
		<Stack>
			{view &&
				<Image boxSize="300px" src={view} alt="select picture" />
			}
			<Stack>
				<Input m="10" name="file" type='file' accept="image/*" onChange={onFileInputChange} />
			</Stack>
		</Stack>
		<Stack>
			<Button m="10" onClick={ApiFetch}　>アイコンを変更する</Button>
			{posts &&
				<Image boxSize="300px" src={(window.URL.createObjectURL(posts))} alt="select picture" />
			}
		</Stack>
		{user && <Stack>
			<FormControl onSubmit={handleSubmit(setData)}>
				<FormLabel>名前を変更する</FormLabel>
				<Input defaultValue={user.Name} {...register("name")} />
			</FormControl>
			<FormControl onSubmit={handleSubmit(setData)}>
				<FormLabel>自己紹介</FormLabel>
				<Input defaultValue={user.Profile.Essay} {...register("essay", {
					pattern: {
						value: /^[^^＾"”'’;； 　#$%<>`*]+$/,
						message: '特殊文字は使えません' // JS only: <p>error message</p> TS only support string
					}
				})}  />
				{errors.essay && errors.essay.message}
				<FormHelperText>200文字以内でお願いします</FormHelperText>
			</FormControl>
			{user.Profile.ID != 0 &&
				<><FormControl onSubmit={handleSubmit(setData)}>
					<FormLabel>Twitter URl</FormLabel>
					<InputGroup>
						<InputLeftAddon>https://twitter.com/</InputLeftAddon>
						{user.Profile.URLs && <Input defaultValue={user.Profile.URLs[0].URL} {...register("url1")} />}
						{!user.Profile.URLs && <Input {...register("url1")} />}
					</InputGroup>
				</FormControl>
					<FormControl onSubmit={handleSubmit(setData)}>
						<FormLabel>Github URl</FormLabel>
						<InputGroup>
							<InputLeftAddon>https://github.com/</InputLeftAddon>
							{user.Profile.URLs && <Input defaultValue={user.Profile.URLs[1].URL} {...register("url2")} />}
							{!user.Profile.URLs && <Input {...register("url2")} />}
						</InputGroup>
					</FormControl></>}
		</Stack>}
		<Spacer />
		<Stack m="10">
			<Button onClick={updateUser} >アカウント情報を更新する</Button>
			<Button onClick={DeleteFetch} disabled={user.ID == guestuser}>ユーザーを削除する</Button>
		</Stack>
	</>)
}
export default Config