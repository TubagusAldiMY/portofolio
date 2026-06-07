import { apiFetch } from './client';
import type {
	ApiDataResponse,
	ApiMessageResponse,
	ChatRequest,
	ChatResponse,
	ContactMessage,
	ContactRequest,
	Experience,
	LoginRequest,
	LoginResponse,
	Product,
	Project
} from './types';

export async function login(payload: LoginRequest): Promise<LoginResponse> {
	return apiFetch<LoginResponse>('/api/auth/login', {
		auth: false,
		body: payload,
		method: 'POST'
	});
}

export async function listProjects(fetcher?: typeof fetch): Promise<Project[]> {
	const response = await apiFetch<ApiDataResponse<Project[]>>('/api/projects', {
		auth: false,
		fetcher
	});
	return response.data;
}

export async function getProject(id: number, fetcher?: typeof fetch): Promise<Project> {
	const response = await apiFetch<ApiDataResponse<Project>>(`/api/projects/${id}`, {
		auth: false,
		fetcher
	});
	return response.data;
}

export async function listProducts(fetcher?: typeof fetch): Promise<Product[]> {
	const response = await apiFetch<ApiDataResponse<Product[]>>('/api/products', {
		auth: false,
		fetcher
	});
	return response.data;
}

export async function listExperiences(fetcher?: typeof fetch): Promise<Experience[]> {
	const response = await apiFetch<ApiDataResponse<Experience[]>>('/api/experiences', {
		auth: false,
		fetcher
	});
	return response.data;
}

export async function sendContactMessage(
	payload: ContactRequest
): Promise<ApiMessageResponse & { data: ContactMessage }> {
	return apiFetch<ApiMessageResponse & { data: ContactMessage }>('/api/contact', {
		auth: false,
		body: payload,
		method: 'POST'
	});
}

export async function sendChatMessage(payload: ChatRequest): Promise<ChatResponse> {
	return apiFetch<ChatResponse>('/api/chat', {
		auth: false,
		body: payload,
		method: 'POST'
	});
}
