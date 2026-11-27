import type { ApiEnvelope } from "./types";

/**
 * Resolve API base URL cleanly for both Next.js SSR (inside Docker network)
 * and Browser/LAN clients.
 */
export function getApiBaseUrl(): string {
  if (typeof window === "undefined") {
    return (
      process.env.INTERNAL_API_URL ??
      process.env.NEXT_PUBLIC_API_URL ??
      "http://localhost:8080/api/v1"
    );
  }
  return process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";
}

export async function apiFetch<T>(
  path: string,
  options?: RequestInit & { idempotencyKey?: string }
): Promise<ApiEnvelope<T>> {
  const headers = new Headers(options?.headers);
  headers.set("Content-Type", "application/json");
  if (options?.idempotencyKey) {
    headers.set("Idempotency-Key", options.idempotencyKey);
  }

  const baseUrl = getApiBaseUrl().replace(/\/+$/, "");
  const normalizedPath = path.startsWith("/") ? path : `/${path}`;

  const res = await fetch(`${baseUrl}${normalizedPath}`, {
    ...options,
    headers,
    cache: "no-store",
  });

  return (await res.json()) as ApiEnvelope<T>;
}

/**
 * Format ISO timestamp into WIB (Asia/Jakarta) per URS 2.1.
 */
export function formatWIB(iso?: string | null): string {
  if (!iso) return "Belum tersedia";
  return (
    new Intl.DateTimeFormat("id-ID", {
      timeZone: "Asia/Jakarta",
      dateStyle: "medium",
      timeStyle: "short",
    }).format(new Date(iso)) + " WIB"
  );
}
