# Tasks: 用户登录功能

**Input**: Design documents from `/specs/001-user-login/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [ ] T001 [P] Create project structure per implementation plan
- [ ] T002 [P] Initialize Go project with necessary dependencies
- [ ] T003 [P] Configure Go modules and dependencies

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T004 Setup PostgreSQL database connection in `internal/repository/db.go`
- [ ] T005 [P] Create password utilities in `pkg/utils/password.go`
- [ ] T006 [P] Create user model in `internal/model/user.go`
- [ ] T007 [P] Create user repository interface in `internal/repository/user_repository.go`
- [ ] T008 [P] Implement PostgreSQL user repository in `internal/repository/postgres_user_repository.go`
- [ ] T009 [P] Setup error handling and logging infrastructure

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - 用户登录 (Priority: P1) 🎯 MVP

**Goal**: 实现用户登录功能，支持用户名和密码验证，登录成功跳转到首页，失败显示错误提示

**Independent Test**: 通过输入正确的用户名和密码测试登录成功流程，通过输入错误的用户名或密码测试登录失败流程

### Tests for User Story 1 ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T010 [P] [US1] Unit test for password utilities in `internal/repository/user_repository_test.go`
- [ ] T011 [P] [US1] Unit test for user repository in `internal/repository/postgres_user_repository_test.go`
- [ ] T012 [P] [US1] Integration test for login endpoint in `tests/integration/auth_test.go`

### Implementation for User Story 1

- [ ] T013 [P] [US1] Create auth service interface in `internal/service/auth_service.go`
- [ ] T014 [P] [US1] Implement auth service in `internal/service/auth_service_impl.go`
- [ ] T015 [P] [US1] Create login request/response models in `internal/model/auth.go`
- [ ] T016 [US1] Implement login handler in `internal/handler/auth.go`
- [ ] T017 [US1] Add login endpoint to router in `cmd/app/main.go`
- [ ] T018 [US1] Add validation and error handling for login
- [ ] T019 [US1] Add logging for login operations

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - 邀请用户 (Priority: P2)

**Goal**: 实现邀请用户功能，通过在用户表中添加用户信息来邀请用户

**Independent Test**: 通过在用户表中添加新用户记录，然后使用该用户的凭据测试登录功能

### Tests for User Story 2 ⚠️

- [ ] T020 [P] [US2] Integration test for invite endpoint in `tests/integration/auth_test.go`

### Implementation for User Story 2

- [ ] T021 [P] [US2] Create invite request/response models in `internal/model/auth.go`
- [ ] T022 [US2] Add invite method to auth service in `internal/service/auth_service_impl.go`
- [ ] T023 [US2] Implement invite handler in `internal/handler/auth.go`
- [ ] T024 [US2] Add invite endpoint to router in `cmd/app/main.go`
- [ ] T025 [US2] Add validation and error handling for invite
- [ ] T026 [US2] Add logging for invite operations

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T027 [P] Update quickstart.md with implementation details
- [ ] T028 [P] Update API spec in contracts/api-spec.json
- [ ] T029 [P] Code cleanup and refactoring
- [ ] T030 [P] Additional unit tests in `tests/unit/`
- [ ] T031 [P] Security hardening (rate limiting, secure headers)
- [ ] T032 [P] Run quickstart.md validation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - May integrate with US1 but should be independently testable

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before endpoints
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Models within a story marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1
   - Developer B: User Story 2
3. Stories complete and integrate independently

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Verify tests fail before implementing
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence
