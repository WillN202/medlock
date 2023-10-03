export type Progress = "" | "STARTED" | "FINISHED"

export interface Achievement {
  id: string
  name: string
  progress: Progress
}