package data

import (
    "context"
    v1 "service/api/account/v1"
    "service/app/account/internal/biz"
    "strconv"

    "github.com/go-kratos/kratos/v2/log"
)

type accountRepo struct {
    data *Data
    log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
    return &accountRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

var accountCacheKey = func(id uint64) string {
    return "account_model_" + strconv.FormatUint(id, 10)
}

var accountEMailCacheKey = func(email string) string {
    return "account_email_to_id_" + email
}

var modelToAccount = func(model *Account) *biz.Account {
    return &biz.Account{
        ID:           model.ID,
        UUID:         model.UUID,
        Email:        model.Email,
        Phone:        &biz.TelPhone{TelCode: model.TelCode, Phone: model.Phone},
        PasswordHash: model.Password,
        Status:       model.Status,
    }
}

func (r *accountRepo) CreateEMailAccount(ctx context.Context, account *biz.Account) (uint64, error) {
    model := &Account{
        UUID:     account.UUID,
        Email:    account.Email,
        TelCode:  account.Phone.TelCode,
        Phone:    account.Phone.Phone,
        Password: account.PasswordHash,
        Status:   account.Status,
    }
    err := r.data.DataBase.Create(&model).Error
    if err != nil {
        return 0, err
    }
    err = r.saveAccountModelToCache(ctx, model)
    if err != nil {
        return 0, err
    }
    return model.ID, nil
}

// saveAccountModelToCache 保存账户数据到缓存
func (r *accountRepo) saveAccountModelToCache(ctx context.Context, model *Account) (err error) {
    // 维护邮箱和id关系 EMail: ID
    err = r.data.Cache.SetString(ctx, accountEMailCacheKey(model.Email), strconv.FormatUint(model.ID, 10), 0)
    if err != nil {
        return
    }

    // TODO 维护手机号和id的关系 未完成

    // 保存用户数据到缓存
    err = r.data.Cache.Set(ctx, accountCacheKey(model.ID), &model, 0)
    if err != nil {
        return
    }
    return
}

// deleteAccountModelToCache 删除账户数据到缓存
func (r *accountRepo) deleteAccountModelToCache(ctx context.Context, model *Account) (err error) {
    // 删除邮箱和id关系 EMail: ID
    err = r.data.Cache.Del(ctx, accountEMailCacheKey(model.Email))
    if err != nil {
        return
    }

    // TODO 删除手机号和id的关系 未完成

    // 删除缓存中的用户数据
    err = r.data.Cache.Del(ctx, accountCacheKey(model.ID))
    if err != nil {
        return
    }
    return
}

func (r *accountRepo) GetAccountByID(ctx context.Context, id uint64) (*biz.Account, error) {
    model := &Account{}
    ok, err := r.data.Cache.Get(ctx, accountCacheKey(id), &model)
    if err != nil {
        r.log.Error("") // TODO
    }
    // 缓存不存在key
    if !ok {
        tx := r.data.DataBase.First(&model, id)
        err = tx.Error
        if err != nil {
            return nil, err
        }

        if tx.RowsAffected == 0 {
            return nil, v1.ErrorInternalServerError("用户不存在 %v", id) // TODO 用户不存在 err
        }

        err = r.saveAccountModelToCache(ctx, model)
        if err != nil {
            r.log.Error("") // TODO
        }
    }
    return modelToAccount(model), nil
}

func (r *accountRepo) GetAccountByEMail(ctx context.Context, email string) (*biz.Account, error) {
    v, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
    if err != nil {
        r.log.Error("") // TODO
    }

    if ok {
        var id, err = strconv.ParseUint(v, 10, 64)
        if err != nil {
            return nil, err
        }
        return r.GetAccountByID(ctx, id)
    }

    // 通过邮箱查找账号数据
    model := &Account{}
    tx := r.data.DataBase.First(&model, "email = ?", email)
    err = tx.Error
    if err != nil {
        return nil, err
    }

    if tx.RowsAffected == 0 {
        return nil, v1.ErrorInternalServerError("用户不存在 %v", email) // TODO 用户不存在 err
    }

    err = r.saveAccountModelToCache(ctx, model)
    if err != nil {
        r.log.Error("") // TODO
    }

    return modelToAccount(model), nil
}

func (r *accountRepo) GetAccountByPhone(ctx context.Context, phone *biz.TelPhone) (*biz.Account, error) {
    // TODO 未完成手机号功能
    return nil, nil
}

func (r *accountRepo) ExistAccountEMail(ctx context.Context, email string) (bool, error) {
    _, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
    if err != nil {
        r.log.Error("") // TODO
    }

    if ok {
        return true, nil
    }

    // 通过邮箱查找账号数据
    model := &Account{}
    tx := r.data.DataBase.First(&model, "email = ?", email)
    err = tx.Error
    if err != nil {
        return false, err
    }

    if tx.RowsAffected == 0 {
        return false, nil
    }

    err = r.saveAccountModelToCache(ctx, model)
    if err != nil {
        r.log.Error("") // TODO
    }

    return true, nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, account *biz.Account) error {
    model := &Account{}
    // 获取原来的账户数据
    tx := r.data.DataBase.First(&model, account.ID)
    err := tx.Error
    if err != nil {
        return err
    }

    if tx.RowsAffected == 0 {
        return v1.ErrorInternalServerError("用户不存在 %v", account.ID) // TODO err
    }

    // 替换数据
    model.UUID = account.UUID
    model.TelCode = account.Phone.TelCode
    model.Phone = account.Phone.Phone
    model.Email = account.Email
    model.Status = account.Status
    model.Password = account.PasswordHash
    // 储存到数据库并加入缓存
    err = r.data.DataBase.Save(&model).Error
    if err != nil {
        return err
    }
    err = r.deleteAccountModelToCache(ctx, model)
    if err != nil {
        r.log.Error("") // TODO
    }
    return nil
}
