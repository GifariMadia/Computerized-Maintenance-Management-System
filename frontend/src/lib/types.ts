export type DivisionCode = "MAINTENANCE" | "UTILITY" | "BUILDING" | "SPARE_PART";

export type UserRole =
  | "ADMINISTRATOR"
  | "TECHNICIAN"
  | "INSPECTOR"
  | "SUPERVISOR"
  | "MANAGER"
  | "SPARE_PART_STAFF";

export type AssetStatus =
  | "ACTIVE"
  | "BREAKDOWN"
  | "UNDER_MAINTENANCE"
  | "STANDBY"
  | "DEACTIVATED";

export type WorkOrderStatus =
  | "OPEN"
  | "IN_PROGRESS"
  | "WAITING_PART"
  | "PENDING_VERIFICATION"
  | "REWORK"
  | "CLOSED"
  | "CANCELLED";

export interface ApiMeta {
  timestamp_wib: string;
  data_stale?: boolean;
  stale_message?: string;
}

export interface ApiEnvelope<T> {
  success: boolean;
  data?: T;
  error?: {
    code: string;
    message: string;
    fields?: Record<string, string>;
  };
  meta: ApiMeta;
}
