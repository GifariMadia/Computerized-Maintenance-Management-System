import type { ApiEnvelope } from "./types";

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api/v1";

export async function apiFetch<T>(
  path: string,
  options?: RequestInit & { idempotencyKey?: string }
): Promise<ApiEnvelope<T>> {
  const headers = new Headers(options?.headers);
  headers.set("Content-Type", "application/json");
  if (options?.idempotencyKey) {
    headers.set("Idempotency-Key", options.idempotencyKey);
  }

  const res = await fetch(`${API_BASE_URL}${path}`, {
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
  return new Intl.DateTimeFormat("id-ID", {
    timeZone: "Asia/Jakarta",
    dateStyle: "medium",
    timeStyle: "short",
  }).format(new Date(iso)) + " WIB";
}
