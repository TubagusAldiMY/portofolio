export type ApiDataResponse<T> = {
	data: T;
};

export type ApiMessageResponse = {
	message: string;
};

export type LoginRequest = {
	username: string;
	password: string;
};

export type LoginResponse = {
	token: string;
};

export type Project = {
	id: number;
	title: string;
	description: string;
	imageUrl: string;
	techStack: string[];
	repoUrl: string;
	liveUrl: string;
};

export type Product = {
	id: number;
	title: string;
	description: string;
	price: string;
	imageUrl: string;
	features: string[];
	buyUrl: string;
	tag: string;
};

export type Experience = {
	id: number;
	role: string;
	company: string;
	duration: string;
	description: string;
};

export type ContactRequest = {
	name: string;
	email: string;
	content: string;
};

export type ContactMessage = ContactRequest & {
	id: number;
	created_at: string;
};

export type ChatRequest = {
	message: string;
};

export type ChatResponse = {
	reply: string;
};

export type UploadResponse = {
	url: string;
	filename: string;
};

export type AdminResource = 'projects' | 'products' | 'experiences';

export type AdminResourceMap = {
	projects: Project;
	products: Product;
	experiences: Experience;
};
