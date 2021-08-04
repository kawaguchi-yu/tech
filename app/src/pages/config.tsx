import React, { useState } from "react"
import {
	Button,
	Input,
	Image,
	Stack,
	Spacer,
} from '@chakra-ui/react';
import router from "next/router";
import Template from "./template";
const Config = () => {
	const [view, setview] = useState<string>();
	const [iconData, setIconData] = useState<FormData>();
	const [posts, setPosts] = useState<Blob>();
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
		<Stack>
			{view &&
				<Image boxSize="300px" src={view} alt="select picture" />
			}
			<Stack>
				<Input m="10" name="file" type='file' accept="image/*" onChange={onFileInputChange} />
			</Stack>
		</Stack>
		<Stack>
			<Button m="10" onClick={ApiFetch}>アイコンを変更する</Button>
			{posts &&
				<Image boxSize="300px" src={(window.URL.createObjectURL(posts))} alt="select picture" />
			}
		</Stack>
		<Spacer />
		<Stack m="10">
			<Button onClick={DeleteFetch}>ユーザーを削除する</Button>
		</Stack>
	</>)
}
export default Config