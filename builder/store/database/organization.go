package database

import (
	"context"

	"github.com/uptrace/bun"
	"jingzhi-server/common/types"
)

type OrgStore struct {
	db *DB
}

func NewOrgStore() *OrgStore {
	return &OrgStore{
		db: defaultDB,
	}
}

type Organization struct {
	ID       int64  `bun:",pk,autoincrement" json:"id"`
	Nickname string `bun:"name,notnull" json:"name"`
	// unique name of the organization
	Name        string     `bun:"path,notnull" json:"path"`
	GitPath     string     `bun:",notnull" json:"git_path"`
	Description string     `json:"description"`
	UserID      int64      `bun:",notnull" json:"user_id"`
	Homepage    string     `bun:"" json:"homepage,omitempty"`
	Logo        string     `bun:"" json:"logo,omitempty"`
	Verified    bool       `bun:"" json:"verified"`
	OrgType     string     `bun:"" json:"org_type"`
	User        *User      `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	NamespaceID int64      `bun:",notnull" json:"namespace_id"`
	Namespace   *Namespace `bun:"rel:has-one,join:namespace_id=id" json:"namespace"`
	Industry    string     `bun:"" json:"industry"`
	times
}

func (s *OrgStore) Create(ctx context.Context, org *Organization, namepace *Namespace) (err error) {
	err = s.db.Operator.Core.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err = assertAffectedOneRow(tx.NewInsert().Model(org).Exec(ctx)); err != nil {
			return err
		}
		namepace.NamespaceType = OrgNamespace
		if err = assertAffectedOneRow(tx.NewInsert().Model(namepace).Exec(ctx)); err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *OrgStore) GetUserOwnOrgs(ctx context.Context, username string, req types.SearchOrgReq, supportPagination bool) (orgs []Organization, total int, err error) {
	query := s.db.Operator.Core.
		NewSelect().
		Model(&orgs).
		Relation("User")
	if req.OrgType != "" {
		query = query.Where("organization.org_type = ?", req.OrgType)
	}
	if req.Industry != "" {
		query = query.Where("organization.industry = ?", req.Industry)
	}
	if req.Search != "" {
		query = query.Where("(organization.name ILIKE ? OR organization.description ILIKE ?)", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if username != "" {
		query = query.
			Join("JOIN users AS u ON u.id = organization.user_id").
			Where("u.username =?", username)
	}

	query = query.Order("organization.updated_at DESC")

	total = 0

	if supportPagination {
		total, err = query.Count(ctx)
		if err != nil {
			return
		}
		err = query.Limit(req.PageSize).Offset((req.Page-1)*req.PageSize).Scan(ctx, &orgs)
	} else {
		err = query.Scan(ctx, &orgs)
	}

	return
}

func (s *OrgStore) Update(ctx context.Context, org *Organization) (err error) {
	err = assertAffectedOneRow(s.db.Operator.Core.
		NewUpdate().
		Model(org).
		WherePK().
		Exec(ctx))
	return
}

func (s *OrgStore) Delete(ctx context.Context, path string) (err error) {
	err = s.db.Operator.Core.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err = assertAffectedOneRow(
			tx.NewDelete().
				Model(&Organization{}).
				Where("path = ?", path).
				Exec(ctx)); err != nil {
			return err
		}
		if err = assertAffectedOneRow(
			tx.NewDelete().
				Model(&Namespace{}).
				Where("path = ?", path).
				Exec(ctx)); err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *OrgStore) FindByPath(ctx context.Context, path string) (org Organization, err error) {
	org.Nickname = path
	err = s.db.Operator.Core.
		NewSelect().
		Model(&org).
		Where("path =?", path).
		Scan(ctx)
	return
}

func (s *OrgStore) Exists(ctx context.Context, path string) (exists bool, err error) {
	var org Organization
	exists, err = s.db.Operator.Core.
		NewSelect().
		Model(&org).
		Where("path =?", path).
		Exists(ctx)
	if err != nil {
		return
	}
	return
}

func (s *OrgStore) GetUserBelongOrgs(ctx context.Context, userID int64) (orgs []Organization, err error) {
	err = s.db.Operator.Core.
		NewSelect().
		Model(&orgs).
		Join("join members on members.organization_id = organization.id").
		Where("members.user_id = ?", userID).
		Scan(ctx, &orgs)
	return
}
